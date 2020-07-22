// Copyright 2018 The go-TITprotocol Authors
// This file is part of the go-TITprotocol library.
//
// The go-TITprotocol library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-TITprotocol library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-TITprotocol library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"crypto/ecdsa"
	"github.com/TITprotocol/go-TITprotocol"
	"github.com/TITprotocol/go-TITprotocol/common"
	"github.com/TITprotocol/go-TITprotocol/contracts/masternode/contract"
	"github.com/TITprotocol/go-TITprotocol/core/types"
	"github.com/TITprotocol/go-TITprotocol/core/types/masternode"
	"github.com/TITprotocol/go-TITprotocol/crypto"
	"github.com/TITprotocol/go-TITprotocol/eth/downloader"
	"github.com/TITprotocol/go-TITprotocol/event"
	"github.com/TITprotocol/go-TITprotocol/log"
	"github.com/TITprotocol/go-TITprotocol/p2p"
	"github.com/TITprotocol/go-TITprotocol/p2p/discover"
	"github.com/TITprotocol/go-TITprotocol/params"
	"github.com/TITprotocol/go-TITprotocol/common/math"
)

var (
	ErrUnknownMasternode = errors.New("unknown masternode")
)

type x8 [8]byte

type MasternodeAccount struct {
	index    int
	id       string
	address  common.Address
	isActive bool
}

type MasternodeManager struct {
	// channels for fetcher, syncer, txsyncLoop
	IsMasternode uint32
	srvr         *p2p.Server
	contract     *contract.Contract

	mux *event.TypeMux
	eth *Ethereum

	syncing int32

	mu          sync.RWMutex
	rw          sync.RWMutex
	ID          string
	NodeAccount common.Address
	PrivateKey  *ecdsa.PrivateKey

	masternodeKeys     map[string]*ecdsa.PrivateKey
	masternodeAccounts map[x8]*MasternodeAccount
}

func NewMasternodeManager(eth *Ethereum) (*MasternodeManager, error) {
	contractBackend := NewContractBackend(eth)
	contract, err := contract.NewContract(params.MasterndeContractAddress, contractBackend)
	if err != nil {
		return nil, err
	}
	// Create the masternode manager with its initial settings
	manager := &MasternodeManager{
		eth:                eth,
		contract:           contract,
		masternodeKeys:     make(map[string]*ecdsa.PrivateKey, params.MasternodeKeyCount),
		masternodeAccounts: make(map[x8]*MasternodeAccount, params.MasternodeKeyCount),
		syncing:            0,
	}
	return manager, nil
}

func (self *MasternodeManager) Clear() {
	self.mu.Lock()
	defer self.mu.Unlock()

}

func (self *MasternodeManager) Start(srvr *p2p.Server, mux *event.TypeMux) {
	self.mux = mux
	log.Info("MasternodeManqager start ")
	for i, key := range srvr.Config.MasternodeKeys {
		id8 := self.X8(key)
		id := fmt.Sprintf("%x", id8[:])
		self.masternodeKeys[id] = key
		account := self.newMasternodeAccount(i, id8, id, key)
		self.masternodeAccounts[id8] = account
		self.activeMasternode(id8)
	}
	self.srvr = srvr
	go self.masternodeLoop()
	go self.checkSyncing()
}


func (self *MasternodeManager) SetMinerKey(index int, etherbase common.Address, key *ecdsa.PrivateKey) (bool, string) {
	for id8_, account := range self.masternodeAccounts {
		if account.index == index {
			id_ := fmt.Sprintf("%x", id8_[:])
			id8 := self.X8(key)
			id := fmt.Sprintf("%x", id8[:])

			delete(self.masternodeAccounts, id8_)
			delete(self.masternodeKeys, id_)

			if account.isActive {
				fmt.Println("Note: The active masternode(", id, ") was replaced!")
			}

			account := self.newMasternodeAccount(index, id8, id, key)
			self.masternodeAccounts[id8] = account
			self.masternodeKeys[id] = key
			self.activeMasternode(id8)
			return true, id
		}
	}
	return false, ""
}

func (self *MasternodeManager) newMasternodeAccount(index int, id8 x8, id string, key *ecdsa.PrivateKey) *MasternodeAccount {
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return &MasternodeAccount{
		index:    index,
		id:       id,
		address:  addr,
	}
}

func (self *MasternodeManager) checkSyncing() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&self.syncing, 1)
		case downloader.DoneEvent, downloader.FailedEvent:
			atomic.StoreInt32(&self.syncing, 0)
		}
	}
}

func (self *MasternodeManager) CheckMasternodeId(id string) bool {
	if _, ok := self.masternodeKeys[id]; ok {
		return true
	}
	return false
}


func (self *MasternodeManager) MasternodeList(number *big.Int) ([]string, error) {
	return masternode.GetIdsByBlockNumber(self.contract, number)
}

func (self *MasternodeManager) GetGovernanceContractAddress(number *big.Int) (common.Address, error) {
	return masternode.GetGovernanceAddress(self.contract, number)
}

func (self *MasternodeManager) SignHash(id string, hash []byte) ([]byte, error) {
	// Look up the key to sign with and abort if it cannot be found
	self.rw.RLock()
	defer self.rw.RUnlock()

	if key, ok := self.masternodeKeys[id]; ok {
		// Sign the hash using plain ECDSA operations
		return crypto.Sign(hash, key)
	}

	return nil, ErrUnknownMasternode
}

func (self *MasternodeManager) GetWitnesses() (ids []string) {
	for id, _ := range self.masternodeKeys {
		ids = append(ids, id)
	}
	return ids
}

// X8 returns 8 bytes of ecdsa.PublicKey.X
func (self *MasternodeManager) X8(key *ecdsa.PrivateKey) (id x8) {
	buf := make([]byte, 32)
	math.ReadBits(key.PublicKey.X, buf)
	copy(id[:], buf[:8])
	return id
}

func (self *MasternodeManager) XY(key *ecdsa.PrivateKey) (xy [64]byte) {
	pubkey := key.PublicKey
	math.ReadBits(pubkey.X, xy[:32])
	math.ReadBits(pubkey.Y, xy[32:])
	return xy
}

func (self *MasternodeManager) masternodeLoop() {
	joinCh := make(chan *contract.ContractJoin, 32)
	quitCh := make(chan *contract.ContractQuit, 32)
	joinSub, err1 := self.contract.WatchJoin(nil, joinCh)
	if err1 != nil {
		// TODO: exit
		return
	}
	quitSub, err2 := self.contract.WatchQuit(nil, quitCh)
	if err2 != nil {
		// TODO: exit
		return
	}

	ping := time.NewTimer(600 * time.Second)
	defer ping.Stop()
	ntp := time.NewTimer(time.Second)
	defer ntp.Stop()

	for {
		select {
		case join := <-joinCh:
			if _, ok := self.masternodeAccounts[join.Id]; ok {
				self.activeMasternode(join.Id)
			}
		case quit := <-quitCh:
			if account, ok := self.masternodeAccounts[quit.Id]; ok {
				fmt.Printf("### [%x] Remove masternode! \n", quit.Id)
				account.isActive = false
			}
		case err := <-joinSub.Err():
			joinSub.Unsubscribe()
			fmt.Println("eventJoin err", err.Error())
		case err := <-quitSub.Err():
			quitSub.Unsubscribe()
			fmt.Println("eventQuit err", err.Error())
		case <-ntp.C:
			ntp.Reset(10 * time.Minute)
			go discover.CheckClockDrift()
		case <-ping.C:
			logTime := time.Now().Format("[2006-01-02 15:04:05]")
			ping.Reset(masternode.MASTERNODE_PING_INTERVAL)
			if atomic.LoadInt32(&self.syncing) == 1 {
				fmt.Println(logTime, " syncing...")
				break
			}
			stateDB, _ := self.eth.blockchain.State()
			contractBackend := NewContractBackend(self.eth)
			for id8, account := range self.masternodeAccounts {
				if account.isActive {
					address := account.address
					if stateDB.GetBalance(address).Cmp(big.NewInt(1e+16)) < 0 {
						fmt.Println(logTime, "Expect to deposit 0.01 etz to ", address.String())
						continue
					}
					gasPrice, err := self.eth.APIBackend.gpo.SuggestPrice(context.Background())
					if err != nil {
						fmt.Println("Get gas price error:", err)
						gasPrice = big.NewInt(20e+9)
					}
					msg := ethereum.CallMsg{From: address, To: &params.MasterndeContractAddress}
					gas, err := contractBackend.EstimateGas(context.Background(), msg)
					if err != nil {
						fmt.Println("Get gas error:", err)
						continue
					}
					minPower := new(big.Int).Mul(big.NewInt(int64(gas)), gasPrice)
					// fmt.Println("Gas:", gas, "GasPrice:", gasPrice.String(), "minPower:", minPower.String())
					if stateDB.GetPower(address, self.eth.blockchain.CurrentBlock().Number()).Cmp(minPower) < 0 {
						fmt.Println(logTime, "Insufficient power for ping transaction.", address.Hex(), self.eth.blockchain.CurrentBlock().Number().String(), stateDB.GetPower(address, self.eth.blockchain.CurrentBlock().Number()).String())
						continue
					}
					tx := types.NewTransaction(
						self.eth.txPool.State().GetNonce(address),
						params.MasterndeContractAddress,
						big.NewInt(0),
						gas,
						gasPrice,
						nil,
					)
					signed, err := types.SignTx(tx, types.NewEIP155Signer(self.eth.blockchain.Config().ChainID), self.masternodeKeys[account.id])
					if err != nil {
						fmt.Println(logTime, "SignTx error:", err)
						continue
					}
					if err := self.eth.txPool.AddLocal(signed); err != nil {
						fmt.Println(logTime, "send ping to txpool error:", err)
						continue
					}
					fmt.Printf("%s [%s] Heartbeat\n", logTime, address.String())
				}else{
					self.activeMasternode(id8)
				}
			}
		}
	}
}

func (self *MasternodeManager) activeMasternode(id8 x8) {
	has, err := self.contract.Has(nil, id8)
	if !has || err != nil || self.masternodeAccounts[id8].isActive {
		return
	}
	self.masternodeAccounts[id8].isActive = true
	id := fmt.Sprintf("%x", id8[:])
	if etherbase, ok := self.eth.etherbases[id]; !ok || etherbase == (common.Address{}) {
		info, error := self.contract.GetInfo(nil, id8)
		if info.Account != (common.Address{}) && error == nil {
			self.eth.SetEtherbaseById(id, info.Account)
		} else if info.Account != (common.Address{}) {
			fmt.Println("[MN] GetInfo Error:", id, error)
		}
	}
	// fmt.Printf("### [%x] Become masternode! \n", id8)
}