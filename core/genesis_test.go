// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/TITprotocol/go-TITprotocol/common"
	"github.com/TITprotocol/go-TITprotocol/consensus/ethash"
	"github.com/TITprotocol/go-TITprotocol/core/rawdb"
	"github.com/TITprotocol/go-TITprotocol/core/vm"
	"github.com/TITprotocol/go-TITprotocol/ethdb"
	"github.com/TITprotocol/go-TITprotocol/params"
	"os"
	"bufio"
	"io"
	"fmt"
	"github.com/TITprotocol/go-TITprotocol/crypto"
	"math/rand"
	"time"
)

func TestInitAccounts(t *testing.T) {
	path := "/Users/rolong/etz/TITprotocol-mns/bin/init.data.00"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufReader := bufio.NewReader(file)
	srcAccountCount := 0

	addrs := make(map[common.Hash]*big.Int)
	nonces := make(map[common.Hash]uint64)

	srcAccountBalance := big.NewInt(00)

	for {
		buf, err := myReader(bufReader, 43)
		if err == io.EOF {
			fmt.Println("Import initial objects", srcAccountCount, "DONE")
			break
		}else if err != nil {
			panic(err)
		}
		addrHash := common.BytesToHash(buf[0:32])
		balance := new(big.Int).SetBytes(buf[32:43])

		nonceB, err := myReader(bufReader, 3)
		if err != nil {
			panic(err)
		}
		nonces[addrHash] = uint64(byte2len(nonceB))

		addrs[addrHash] = balance
		srcAccountBalance.Add(srcAccountBalance, balance)
		srcAccountCount++
		if srcAccountCount % 10000 == 0 {
			fmt.Println("Import initial objects", srcAccountCount)
		}
	}

	// add accounts
	addrs[myHash("0x4d7aa619eb0e2ef0a59b4febb09f86c8c5430a96")] = newBalance("9463576726764413066000")
	addrs[myHash("0xed0328c3e6d99cfa20f8206eff274a24c048e352")] = newBalance("9693622222833470400000")
	addrs[myHash("0x05404a5226791fc68d04666e6eae2a7ff2c18d8d")] = newBalance("9923637222091693809000")
	addrs[myHash("0x593962C1CbFBD9Eed0e166f36C9ff2478A2517BA")] = newBalance("61073314913552061000000")
	addrs[myHash("0xa6dbac124c213aa38159b4502a8af35fdcb00a1d")] = newBalance("10153492980794284586000")
	addrs[myHash("0x88ee89a83fe03b32a9a4bf6debddc41f8eb2de97")] = newBalance("10383360410594550343000")
	addrs[myHash("0x62e94144623e8c1dc1aee9f084f859bcd74f9381")] = newBalance("10613726131522525649000")
	addrs[myHash("0x6443fd3fbe031ba0e7eecf29e599dc2fbece3e2c")] = newBalance("10843690740841243711000")
	addrs[myHash("0x2637c14ade8c96fdcf4eb3cf0e30079c0c61a5f6")] = newBalance("11073377707145105101000")
	addrs[myHash("0x83ac31c9bbc7ca4b94ab7a93e0ccd356d7cae060")] = newBalance("8303819390248428303000")
	addrs[myHash("0x0d33bf9d791b2638c3c360bbdb5d4a43f77866c1")] = newBalance("11533309449069955056000")
	addrs[myHash("0xdc29fbf6bd28a4132372af12b696bdc88a5ef94a")] = newBalance("7763556614527647663000")
	addrs[myHash("0x55e8f670205eee019679ff8ae8e580743a565fa8")] = newBalance("11993226093581465961000")
	addrs[myHash("0x894f5360f0874b64473b23faea1e871e2b935423")] = newBalance("12223509803293317544000")
	addrs[myHash("0x313a8744c5643eb7253cf3470e91227a22c99c77")] = newBalance("12453411827295321008000")
	addrs[myHash("0x02a9197606ea987dad07768566492f263561ab77")] = newBalance("9683496517616215310000")
	addrs[myHash("0x456aaf288b925ce69f057b2498e09099962e9700")] = newBalance("12913461830686347194000")
	addrs[myHash("0x671302d355efb628f4197d68d971e459c8a64271")] = newBalance("13143499621675458701000")
	addrs[myHash("0x9579395beb243d79728b16db6dbbbe2bb73fa1fe")] = newBalance("13373245815249799888000")
	addrs[myHash("0x90cc248600490f4b97ebb4823adf5dd16aac0fb4")] = newBalance("13603447060212347660000")
	addrs[myHash("0xfaacc85495959e73351c602a546979e0db9258f7")] = newBalance("13833522861616890962000")
	addrs[myHash("0xbd6660d3db91432e521ab3ab84e88e1d0ebafcb9")] = newBalance("14063774007685193980000")
	addrs[myHash("0x426e204185c131fb375a3496ae34bbecd625e749")] = newBalance("14293598330821591598000")
	addrs[myHash("0xadf5f1c5214acfef911102d831e70f91b0fd21a0")] = newBalance("14523323403223359296000")
	addrs[myHash("0x408e3175325239789e3da2f322e59b94f44cb47a")] = newBalance("14753532824709187459000")
	addrs[myHash("0xe27511e3cb75076d51d6bfc503419b081e2faf2c")] = newBalance("14983416580917527336000")
	addrs[myHash("0xb3a5ad5a571ac31a04e2c9cc5f45380218b832bd")] = newBalance("8213193140648761067000")
	addrs[myHash("0x382212f58af48a263ac7c540e080bec1dd64bf71")] = newBalance("15443149253117194275000")
	addrs[myHash("0x6b0636a80ec75633f8d81f1e26a52334e7f30b49")] = newBalance("15673458088220560417000")
	addrs[myHash("0x4fa72ed959d8a5e5e30b123c9f2e6e9cd70c268c")] = newBalance("15903580418228562192000")
	addrs[myHash("0xa2c11efd91fb2670cfccd3dbcaec2dc1e60a31ee")] = newBalance("16133705764855778631000")
	addrs[myHash("0x3360bdeb4e075c3d762b1995e1cd94787ac50750")] = newBalance("16363725330968485608000")
	addrs[myHash("0x5c1f72a20b613fa79a6911d51947a4eb33d3aa6e")] = newBalance("16593816216722875297000")
	addrs[myHash("0x0e97bbb77f845d61e0bd2e68746acac00c6ff526")] = newBalance("16823419915444218670000")
	addrs[myHash("0xbfca1421c4d0e938b73c3ab19a3816739cc819f4")] = newBalance("17053671795147817430000")
	addrs[myHash("0x05acd53577a45b9779306a00a1dbcc9421b42d98")] = newBalance("17283513681291733940000")
	addrs[myHash("0xc6a92462462bc8d0a044ad7ab7771b04c7f10677")] = newBalance("17513906546737985630000")
	addrs[myHash("0x3a599950435265ff3f7ed3de7a48aa1408d3fad1")] = newBalance("17743672821612787250000")
	addrs[myHash("0xbac6f7e61cba2c1614e3f9effae422e2e3f3e078")] = newBalance("17973475199206697180000")
	addrs[myHash("0xb2e44bab1e68e7f1cfbcd0b6c74d44fcc962ff54")] = newBalance("18203301286675833190000")
	addrs[myHash("0xe6ef9179211ba9d7515bc8dd8491b80702aecd8c")] = newBalance("18433528910423455620000")
	addrs[myHash("0x21ac18e704e3994f19178bde5af8faa5cd3800e5")] = newBalance("18663865595694920530000")
	addrs[myHash("0x8e75183ecc9e79c47fd6d5f64382563c13047275")] = newBalance("18893569603819880760000")
	addrs[myHash("0x0428b3da274a5eed3a6f28f3821dcca664371e51")] = newBalance("19123483846518948090000")
	addrs[myHash("0xd4e09ebe005e33a711276392540110c6e7b9d94d")] = newBalance("19353610148251859690000")
	addrs[myHash("0xbf605e4de1a9dc121dde6ad66eccf80273bd5ded")] = newBalance("6583843331979636870000")
	addrs[myHash("0xfcaf8b433bb25834cc2c59be78a0aebc0d625e3b")] = newBalance("19813598542757454320000")
	addrs[myHash("0xe139f55ab67cd34c73c31166afc9cbae9e6e1fd3")] = newBalance("20043445799171973690000")
	addrs[myHash("0x7d4a6e807e5133e270d8a89ba99c49ee0a963517")] = newBalance("20273576474846857970000")
	addrs[myHash("0xa9943cdc860c4b1004f888521a320f1de13a6802")] = newBalance("20503516304492914370000")
	addrs[myHash("0x92f8746dbe554b772a1429a7ab91a83a40c4088f")] = newBalance("20733509325844618090000")
	addrs[myHash("0x957aeba053bc564e4d1a1ba6fd958364b26f45aa")] = newBalance("722180445631686865220000")
	addrs[myHash("0xf255a98d07851f34fb6a52ddfee9b792248fca75")] = newBalance("722083646842147556720000")
	addrs[myHash("0x50271f71f3d5f339b6562f3ffba8fb3b8bab5858")] = newBalance("708120329055232433560000")
	addrs[myHash("0x86ffdda609a2077cdd4db24935cd8c731aee3c69")] = newBalance("706583410280858899020000")
	addrs[myHash("0x11a949304880f652156e9f04eaa31f0b74a260f4")] = newBalance("706342388506633357040000")
	addrs[myHash("0xdbe5f1f7cb4245a7e66caa988571d9c11c5bf43b")] = newBalance("706245157662266824000000")
	addrs[myHash("0x94f0bf41622219c81f3afb2da48a93a87b87acc6")] = newBalance("706148877922051855000000")
	addrs[myHash("0x77f9aae0fb63b35a1455a2b317db7fcbad33cc1a")] = newBalance("691562645123802246000000")
	addrs[myHash("0xf38bd669eddf6f76bc658127bb97e1c682936ab8")] = newBalance("671749123454535874000000")
	addrs[myHash("0x31b830973de5321b0f81ea88a5f03c5e478e1f17")] = newBalance("671652373726439218000000")
	addrs[myHash("0xa2df17de25e1dd672248a6ebe95d4df910593a54")] = newBalance("666902531536491605000000")
	addrs[myHash("0x29811ca2eb86229db6ec1fdc1a51cf7228ce2d5f")] = newBalance("666805460603340185000000")
	addrs[myHash("0x7a0ab70548de239b368ab5fef4375a0fc10eac69")] = newBalance("665550383816328201000000")
	addrs[myHash("0x889e8c493c174dfbbbfa5853cb44166300025668")] = newBalance("664064974758476745000000")
	addrs[myHash("0xb784b3f1c379df5773dd826ca8e4fca616af9eda")] = newBalance("663967822129100878000000")
	addrs[myHash("0xd452725cc0af5402c3d0731fca95751764491cda")] = newBalance("663786125256315857000000")
	addrs[myHash("0x12374c1c2521350029e1b348f97dd969cb1ebf93")] = newBalance("663278411062724418000000")
	addrs[myHash("0x88e60e8c995b21a47c2a1de901e6c5477e72b6ee")] = newBalance("663181175396684831000000")
	addrs[myHash("0x67e68d9e9482c5e79a8bd5a21d5f6bf67adedbed")] = newBalance("657753148666572673000000")
	addrs[myHash("0xb4be8dd2be1f72174bed560a7d34a7c56ce8a71f")] = newBalance("656995420893244582000000")
	addrs[myHash("0x3ddd61c5122ec0a1b71d2b6cdd7eb33d8bcfc440")] = newBalance("649986121363541683000000")
	addrs[myHash("0x43f4b22eef90e365954abc4aeb487f07ff220d8d")] = newBalance("649088398953426128000000")
	addrs[myHash("0x4025e6adcb0aeae6044fbbe277dd156df9af448a")] = newBalance("648991665000555815000000")
	addrs[myHash("0x33b6c53b32c2115f53dcac8c9e81f0a0ce14dbae")] = newBalance("645735587217091679000000")
	addrs[myHash("0xe1e54aec7ae571272be9e7d1155420e8acbe234f")] = newBalance("645530906147514114000000")
	addrs[myHash("0x0c067fc2527ada0e31e9be58a874fa2f52f468a2")] = newBalance("634289422264377736000000")
	addrs[myHash("0x7f48a67dbbd6754cf133561636e67c0ff750dc56")] = newBalance("632217601223255399000000")
	addrs[myHash("0x390e83be240711fa3ef93ad53a7784245a3d107e")] = newBalance("631950930017646453000000")
	addrs[myHash("0x7f38f47c45cedb19cc0e7704b00973ad6f3b91e9")] = newBalance("631853148566295769000000")
	addrs[myHash("0x7ec6314ef38bdc20edba6f8f846842ec2bef9ddb")] = newBalance("629181456713790850000000")
	addrs[myHash("0x7da7911dd61c07d9fb436fccf1284c824abefd5e")] = newBalance("616256549465373367000000")
	addrs[myHash("0x1674fdb21ffa908ae4bb4c1538512761a3cb590b")] = newBalance("615365382189124448373000")
	addrs[myHash("0xdc4d212ea32a8c54fde2ac346534c5a8a6259370")] = newBalance("615268472023594393998000")
	addrs[myHash("0xc959bf7f1d8702594d0bbf8158136fd7f0a6a9d1")] = newBalance("615171258460357133385000")
	addrs[myHash("0x0ed29e9650722669d3a5caf055b874ab5835f58f")] = newBalance("615074284361973243640000")
	addrs[myHash("0xfe23c54636b89067c7dea51e47abb883356e4b02")] = newBalance("614763239249939353404000")
	addrs[myHash("0xfe0642afad78ed98c312e4ca3b9977c7baaa85cb")] = newBalance("609222682271444553097000")
	addrs[myHash("0xcfb3c4f5d5388fea1593435dfc9b2b00b62fefce")] = newBalance("606125840645459253690000")
	addrs[myHash("0xdffff59d3e5dcfe5a06dc4fd0e2b34976ec3e0b5")] = newBalance("605853284429377803446000")
	addrs[myHash("0xd3d5ec8fdeb482f058bf9660941b58acbc2e4381")] = newBalance("605459363532624509242000")
	addrs[myHash("0xe5226860fb2ee5b2272cfc96e9d636b2755dcd4b")] = newBalance("603434157678859871414000")
	addrs[myHash("0xf96da875f9d299375f7f0c42ba2821ee420f84e5")] = newBalance("602797396299963152383000")
	addrs[myHash("0x1fe745c913403f4e7d7825e1a186a0690244a89c")] = newBalance("602700621704206588159000")
	addrs[myHash("0xb7d6a40169b90799ff50f90660adc1d79c275f34")] = newBalance("602573925843424208529000")
	addrs[myHash("0xc474cb8a8662e4a75b0cef6bd892695ea4a13b4a")] = newBalance("602476299982374287235000")
	addrs[myHash("0xfddb0102a2a56a5dca4bd5d786285a6418d45240")] = newBalance("586023524550782317826000")
	addrs[myHash("0x182092a7d3a1cc0682e70ec4d171c3857d80feb8")] = newBalance("577088489619755650975000")
	addrs[myHash("0x3dfe096df28e4b8f3921b4d6d7ae093376baf7de")] = newBalance("552155382088470172229000")
	addrs[myHash("0x2e4f23534d977ea751793b5d171ac05a91b295f4")] = newBalance("538806389916865505794000")
	addrs[myHash("0x8e9b58a89a0bdd3f49fd1da1a718bd90dce48357")] = newBalance("537961841168531245987000")
	addrs[myHash("0x29f64ec439b8f76c7cdc54c66a034d93f3dda52d")] = newBalance("531104909087838647825000")
	addrs[myHash("0xca2e7f2b03e0948636c1058e4fba4eb373c3186f")] = newBalance("506767335703134618186000")
	addrs[myHash("0xd41c2f5fc73a8aa29aa51ac90634375f311c9c6e")] = newBalance("505270929800720663695000")
	addrs[myHash("0xb88dad959e6347919f0f5e327111402616bb726a")] = newBalance("504497113572578667894000")
	addrs[myHash("0xc3a806be00333838a81968cd9556ad2b30c7c643")] = newBalance("501526247779894450272000")
	addrs[myHash("0x4a97e36cadc7c981b8d29980498a7a7220534d6c")] = newBalance("501415803532246294917000")
	addrs[myHash("0xc7b7919c246a6ae0045dafc9e5891942c594f657")] = newBalance("491488378938560458246000")
	addrs[myHash("0x965d587d138fe978b3a3c40a9a3e6463acbcaa55")] = newBalance("481185331344683894750000")
	addrs[myHash("0xad74c3cc9e57c68a2f3a535273ce0d3a07115868")] = newBalance("476091280085421784435000")
	addrs[myHash("0x051a8e976204f97110d6e45387cbce152da2aa3b")] = newBalance("470011165197179412082000")
	addrs[myHash("0x27cf64d6d1ff34723dc39dc4c4e33484ffb207cc")] = newBalance("192684519883343192981000")
	addrs[myHash("0xf7d7313ebcd991ecda74d827437548debc35c5c9")] = newBalance("270187578843606812379000")
	addrs[myHash("0xe570d5023a9fae4420b82c9bb741f72efc1b943a")] = newBalance("129470234838243926338000")
	addrs[myHash("0x5e79f08d84a317d68a1fec6e03d83d3fcf8f8810")] = newBalance("124091872610258611005000")
	addrs[myHash("0x22c24939bbb0adcaa27da0ac7175f48fb51cdbc8")] = newBalance("318073161755979488242000")
	addrs[myHash("0x9637c6a79c9779127ad350a0e2a1a215d78366f8")] = newBalance("117219370693372858151000")
	addrs[myHash("0x89e3df43b6a0c09f60c1f85122a79b4d1fbb4bb2")] = newBalance("895151309393218528124000")
	addrs[myHash("0x5bb80660ab0478b0abdc5a9c020d7c0ef938c7c5")] = newBalance("890983723960492155115000")
	addrs[myHash("0x3fc40d9e3266dfdcb5d2a983fa4b6cdeb57dcd8b")] = newBalance("882900797964142140423000")
	addrs[myHash("0x91b0615ba9fc4233dc85ab60b969f33248e549fe")] = newBalance("880724567358500127595000")
	addrs[myHash("0x2154b70cbe5b7be059bd95c0132753f4cb01f38b")] = newBalance("871486960500822509292000")
	addrs[myHash("0xddccde2eca0036f348eae7a94374841382ee068b")] = newBalance("867927134775733746848000")
	addrs[myHash("0x0afff001f9cf599057569c80785b491d50c9a214")] = newBalance("863552672102214830630000")
	addrs[myHash("0x62a86a42ddc99c76998af43e50ebe94abaf7a57b")] = newBalance("859586909839390382975000")
	addrs[myHash("0xf43144a008576d3357f60e7e36cfa70767c55a75")] = newBalance("852173477394526478650000")
	addrs[myHash("0xa2dffb77612f42cf417830411c76855eaee5a47d")] = newBalance("831919114422515464418000")
	addrs[myHash("0xb62b11bfee9992da72f854e9a24d3b4e189c0a57")] = newBalance("824175561326807732776000")
	addrs[myHash("0x9788005e5aa0036eefa5d4d93722df6f7f4b4a71")] = newBalance("818705143068682997368000")
	addrs[myHash("0x012b89f1ad911cb2ec61ca55851eccabdc8c2282")] = newBalance("811381832996431123800000")
	addrs[myHash("0x4d879c80ab44502db6e36c3c9f1aa708cbf4a3ea")] = newBalance("810048508024269818345000")
	addrs[myHash("0x99638aaafddfbf9a80619140eb2e4a4de56bfb30")] = newBalance("808000215221779947395000")
	addrs[myHash("0x6d4edb87c69dcc2cbfe191ffdbde859606dd6d4f")] = newBalance("805298908912756404814000")
	addrs[myHash("0x7ccc3fc38b2a8793133c3b19d4a567d67ec77688")] = newBalance("804706297780559338923000")
	addrs[myHash("0x9592dd68d52c299f8df6f1c7cee471f6c939ce99")] = newBalance("787915401891053909644000")
	addrs[myHash("0x225c4d1c0e81ae70a9992208830bdec269accf03")] = newBalance("787818498804363934922000")
	addrs[myHash("0x817608d3e84156e7bb10a96f3d8d8423b3a97821")] = newBalance("782033925931022820509000")
	addrs[myHash("0xb34e1f713393fa8788c9015f13964b08b4894788")] = newBalance("781936552695141246460000")
	addrs[myHash("0x3a7b4f155efff21a2a7c4a9974887c21aaebaf23")] = newBalance("776037878388909300619000")
	addrs[myHash("0xee90f1f7b3e48ab649f0849b7b70e70cef265320")] = newBalance("775153186220608565795000")
	addrs[myHash("0xce67bef7aba7f8bd50a32abfbd275d0a9efbb897")] = newBalance("775015204819721235918000")
	addrs[myHash("0x5449178905822450b0292c9a828a8e7ea8f1323b")] = newBalance("774918452374521746003000")
	addrs[myHash("0x7b08a38fa5daeabd67cfa9d9c320dd584eee2e4d")] = newBalance("774821886165723712551000")
	addrs[myHash("0xffa3a1254da1f8595db6b16d716b3830dd3ceb77")] = newBalance("774724954413655668521000")
	addrs[myHash("0x4007abe3c180bf11e2cac0ec16d799eb1a0d94c5")] = newBalance("763134789056607251727000")
	addrs[myHash("0xb2fef2740f0ecd6f2066f316fb1e736219be1e91")] = newBalance("763037679280669718221000")
	addrs[myHash("0xf613ad937aaa3b16a27f6ab5ad0d25b8f6745b93")] = newBalance("762445594348573101780000")
	addrs[myHash("0x4431e3b3a1b3bdcc1938f639fe1b3e3b46330a24")] = newBalance("762348662652005865551000")
	addrs[myHash("0x64f7f83db71282a7fe50df20341626d5bc74e186")] = newBalance("755957705676748501156000")
	addrs[myHash("0xf909a23268018b0c2a860ef312588d173abdfa7d")] = newBalance("755860115384022709917000")
	addrs[myHash("0xe52461b6286e7c7c48412710cf76a8dd85ae8e91")] = newBalance("755169505533446937107000")
	addrs[myHash("0x7aa492a2a6e5370a53eff34bd58015021ca20893")] = newBalance("755072750618651494813000")
	addrs[myHash("0xd49de50b5bf059777956bd1d0d60d4dd0cfc649b")] = newBalance("737563766129673396753000")
	addrs[myHash("0xeccb556b5fdbcd9437a29620fe0c4224d6bd4b2a")] = newBalance("734298331920509373407000")
	addrs[myHash("0x9443756920bd3a6cc50a8fb31c41251b3c24b854")] = newBalance("727069178330829218756000")
	addrs[myHash("0xb776e79315a490f2488ca73a33398340eaa9868a")] = newBalance("724101880028367569848000")
	addrs[myHash("0xe179c30729c50458a3ba68dabdcb5710f70a3789")] = newBalance("723963461718528399698000")
	addrs[myHash("0x54f13030a1a5cbb6a46841d163103c1697fa0f11")] = newBalance("723866826746953714595000")
	addrs[myHash("0x5b504b9c41e5018ecdd6ef15841e5f1da6f10fdf")] = newBalance("723115947300945821050000")
	addrs[myHash("0x91962ef9140b045ce0954b0d9387801041662ad9")] = newBalance("723018107442000103419000")
	addrs[myHash("0x3d000d35315f3e180304dc2ee28a572d4ce416cd")] = newBalance("424515923934260297533000")
	addrs[myHash("0xd74528c267d0c0f2c24cd9cfc74de337ebbf5c99")] = newBalance("419516879369225970705000")
	addrs[myHash("0xd42b0608185abf53d5f66d811a651ac6b1fee5d7")] = newBalance("414517724037086433505000")
	addrs[myHash("0x56e8eeeb6f133c4f164c62c59e57af183f8ec4cd")] = newBalance("409518246728488816603000")
	addrs[myHash("0x45245412be04d17d6e7edcd2294995bf714f2653")] = newBalance("404519453881916448485000")
	addrs[myHash("0x803adb8db5a11033703e1445dc2d51cd0e103df9")] = newBalance("399520212854204502205000")
	addrs[myHash("0xe7106dc3fbff32ca0ae95663e0f6239cda2f2fa7")] = newBalance("394521493919067477828000")
	addrs[myHash("0x9cf3243cb9692a8b525a9f9fb93f158628959db2")] = newBalance("389522898536704675741000")
	addrs[myHash("0xef97156c5cd180be03d39ef21b79cf39684e31b6")] = newBalance("384523506096105190954000")
	addrs[myHash("0x6f59c0356785d18bb834dcc169300602fe27b770")] = newBalance("379524487528553516761000")
	addrs[myHash("0x0a4db37e010671a858a24653019dd93564b0e051")] = newBalance("374525286645740666427000")
	addrs[myHash("0xea22fb601b32e743808f454d6633b4260824376d")] = newBalance("369526103851257508188000")
	addrs[myHash("0x9269ab9b1dfb8de0c147de82377b0672fe341bde")] = newBalance("364527269551078234245000")
	addrs[myHash("0x3c5ad92173303d6ba9128927b7e334adaefcf25c")] = newBalance("359528732219360973999000")
	addrs[myHash("0xc38978a071b099f0e77108f87832dcc81a01f27b")] = newBalance("354529491621057193249000")
	addrs[myHash("0x20617fbd45c30c7ee50a013208d98fad1f9c9b9f")] = newBalance("349530895832701608954000")
	addrs[myHash("0xd7a35167647ff74825f1f533a4d15f24309a06d7")] = newBalance("344531678283966726928000")
	addrs[myHash("0xfbee69c222ed2e4df3f763c0d36ba5ec3477c9d5")] = newBalance("339532962633082847785000")
	addrs[myHash("0x97a51a187405b815e0bc3998ecfd846a5618ca15")] = newBalance("334533171009009196058000")
	addrs[myHash("0xf63c531a2381f885ccbeb87736117f8a511134a3")] = newBalance("329534711284402478843000")
	addrs[myHash("0x340ced0a45cbf66ff1930cb59dfd706287405f05")] = newBalance("324535660398167526108000")
	addrs[myHash("0xe8edf29c73d3d8a55d790086b6f3bc323c9e40ba")] = newBalance("319536844686109302410000")
	addrs[myHash("0xccc96e126f17b7476f17b5208a1d2479f5761c67")] = newBalance("314537855527662513177000")
	addrs[myHash("0xbc976220d0e7ae1623064898dbb3ac7755b0e407")] = newBalance("309538825159681834302000")
	addrs[myHash("0x86fe35db52f9be4d30a0b7b215d8b04150c41579")] = newBalance("304539685754095144605000")
	addrs[myHash("0x787be04194df2ddc47ccf3f9778a2395c3ed9439")] = newBalance("299540652822911689380000")
	addrs[myHash("0x563e2ef58db64c347a09882ea7a335c2ed19e5db")] = newBalance("294541205118246809838000")
	addrs[myHash("0x49431fa8779bb20402915f53495360fe0180122b")] = newBalance("289542820978422522494000")
	addrs[myHash("0xe8dfe3ba918bba48c8ba71873dbdf2d9a099d739")] = newBalance("284543807972733530544000")
	addrs[myHash("0x1ebad0cf295e9817ece7ee879a0d20d4422d2da9")] = newBalance("579544875924139941012000")
	addrs[myHash("0x8aba980ba99e1f45e6ebff206975d7a568d3297d")] = newBalance("474545922066771916485000")
	addrs[myHash("0xf672c01d56e920b5b75d068781d80aa27e491423")] = newBalance("169546127255065381490000")
	addrs[myHash("0x626dce6c345133ef0995b6190c662d960f06ab86")] = newBalance("264547489339092447050000")
	addrs[myHash("0xe052449f15c216629f1548725a67d08525903832")] = newBalance("259548561265280189918000")
	addrs[myHash("0xd7f985655705e51b2525a8c81757400f4d925427")] = newBalance("614309356930886301867000")
	addrs[myHash("0x6f00141ce741714a5c080229f3887afdd2f3097a")] = newBalance("799350827931578150664000")
	addrs[myHash("0x09c61a45c4a578c9972afc38c29c91885523a734")] = newBalance("384391231878397942246000")
	addrs[myHash("0x35eb651fcef1294a9f7a178c7ba4bddd007eb026")] = newBalance("569432351314212901283000")
	addrs[myHash("0x5d2c72a06676f006c9e846fc58c38bcebb14bbe3")] = newBalance("454473156444930514101000")
	addrs[myHash("0xf59953a1b4f6974b8539968d3edfe5d4526894e9")] = newBalance("339514772144861513785000")
	addrs[myHash("0xe453db4bf302a61e2b31306344eba3b83da11db7")] = newBalance("310555889963488722077000")
	addrs[myHash("0x4985090a0bb648c0057ce89122cc496ed65e536c")] = newBalance("209373502628360429254000")

	// write ===========

	dstAccountBalance := big.NewInt(0)

	dstPath := "/Users/rolong/etz/TITprotocol-mns/bin/init.data.0-test"
	f, err := os.Create(dstPath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	dstAccountCount := 0

	for k, v := range addrs {
		if v.Cmp(common.Big0) <= 0 {
			if v.Cmp(common.Big0) < 0 {
				panic("error balance")
			}
			continue
		}
		bin := append(k[:], common.LeftPadBytes(v.Bytes(), 11)...)
		if len(bin) != 43 {
			panic("error len")
		}
		w.Write(bin)

		nonce, ok := nonces[k]
		if !ok {
			rand.Seed(time.Now().UnixNano())
			nonce = uint64(rand.Intn(1000))
		}
		nonceBytes := nonce2bytes(int(nonce))
		w.Write(nonceBytes)
		dstAccountBalance.Add(dstAccountBalance, v)
		dstAccountCount++
	}

	targetBalance, _ := new(big.Int).SetString("182578076411834954237581102", 10)
	diffBalance := new(big.Int).Sub(targetBalance, dstAccountBalance)
	if diffBalance.Cmp(common.Big0) > 0 {
		fmt.Println("DiffBalance:", diffBalance.String())
		fmt.Println("DiffBalance:", new(big.Int).Div(diffBalance, big.NewInt(params.Ether)).String())
		bin := append(myHash("0x9ba2d229e9d5623284f4750c0d4b6fcd52256b38").Bytes(), common.LeftPadBytes(diffBalance.Bytes(), 11)...)
		if len(bin) != 43 {
			panic("error len")
		}
		w.Write(bin)
		rand.Seed(time.Now().UnixNano())
		nonce := rand.Intn(1000)
		nonceBytes := nonce2bytes(int(nonce))
		w.Write(nonceBytes)
		dstAccountBalance.Add(dstAccountBalance, diffBalance)
		dstAccountCount++
	}

	w.Flush()

	fmt.Println("[ACTIVE3] Src Accounts Count   :", srcAccountCount)
	fmt.Println("[ACTIVE3] Src Accounts Balance :", new(big.Int).Div(srcAccountBalance, big.NewInt(params.Ether)).String())
	fmt.Println("[ACTIVE3] Src Accounts Balance :", srcAccountBalance.String())
	fmt.Println("------")
	fmt.Println("[ACTIVE4] Dst Accounts Count   :", dstAccountCount)
	fmt.Println("[ACTIVE4] Dst Accounts Balance :", new(big.Int).Div(dstAccountBalance, big.NewInt(params.Ether)).String())
	fmt.Println("[ACTIVE4] Dst Accounts Balance :", dstAccountBalance.String())

}

func newBalance(v string) *big.Int {
	balance, status := new(big.Int).SetString(v, 10)
	if !status {
		fmt.Println("newBalance", v)
		panic(v)
	}
	return balance
}

func myHash(v string)  common.Hash {
	return crypto.Keccak256Hash(common.FromHex(v))
}

func TestDefaultGenesisBlock(t *testing.T) {
	block := DefaultGenesisBlock().ToBlock(nil)
	if block.Hash() != params.MainnetGenesisHash {
		t.Errorf("wrong mainnet genesis hash, got %v, want %v", block.Hash(), params.MainnetGenesisHash)
	}
	block = DefaultTestnetGenesisBlock().ToBlock(nil)
	if block.Hash() != params.TestnetGenesisHash {
		t.Errorf("wrong testnet genesis hash, got %v, want %v", block.Hash(), params.TestnetGenesisHash)
	}
}

func TestSetupGenesis(t *testing.T) {
	var (
		customghash = common.HexToHash("0x89c99d90b79719238d2645c7642f2c9295246e80775b38cfd162b696817fbd50")
		customg     = Genesis{
			Config: &params.ChainConfig{HomesteadBlock: big.NewInt(3)},
			Alloc: GenesisAlloc{
				{1}: {Balance: big.NewInt(1), Storage: map[common.Hash]common.Hash{{1}: {1}}},
			},
		}
		oldcustomg = customg
	)
	oldcustomg.Config = &params.ChainConfig{HomesteadBlock: big.NewInt(2)}
	tests := []struct {
		name       string
		fn         func(ethdb.Database) (*params.ChainConfig, common.Hash, error)
		wantConfig *params.ChainConfig
		wantHash   common.Hash
		wantErr    error
	}{
		{
			name: "genesis without ChainConfig",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				return SetupGenesisBlock(db, new(Genesis))
			},
			wantErr:    errGenesisNoConfig,
			wantConfig: params.AllEthashProtocolChanges,
		},
		{
			name: "no block in DB, genesis == nil",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				return SetupGenesisBlock(db, nil)
			},
			wantHash:   params.MainnetGenesisHash,
			wantConfig: params.DevoteChainConfig,
		},
		{
			name: "mainnet block in DB, genesis == nil",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				DefaultGenesisBlock().MustCommit(db)
				return SetupGenesisBlock(db, nil)
			},
			wantHash:   params.MainnetGenesisHash,
			wantConfig: params.DevoteChainConfig,
		},
		{
			name: "custom block in DB, genesis == nil",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				customg.MustCommit(db)
				return SetupGenesisBlock(db, nil)
			},
			wantHash:   customghash,
			wantConfig: customg.Config,
		},
		{
			name: "custom block in DB, genesis == testnet",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				customg.MustCommit(db)
				return SetupGenesisBlock(db, DefaultTestnetGenesisBlock())
			},
			wantErr:    &GenesisMismatchError{Stored: customghash, New: params.TestnetGenesisHash},
			wantHash:   params.TestnetGenesisHash,
			wantConfig: params.TestnetChainConfig,
		},
		{
			name: "compatible config in DB",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				oldcustomg.MustCommit(db)
				return SetupGenesisBlock(db, &customg)
			},
			wantHash:   customghash,
			wantConfig: customg.Config,
		},
		{
			name: "incompatible config in DB",
			fn: func(db ethdb.Database) (*params.ChainConfig, common.Hash, error) {
				// Commit the 'old' genesis block with Homestead transition at #2.
				// Advance to block #4, past the homestead transition block of customg.
				genesis := oldcustomg.MustCommit(db)

				bc, _ := NewBlockChain(db, nil, oldcustomg.Config, ethash.NewFullFaker(), vm.Config{}, nil)
				defer bc.Stop()

				blocks, _ := GenerateChain(oldcustomg.Config, genesis, ethash.NewFaker(), db, 4, nil)
				bc.InsertChain(blocks)
				bc.CurrentBlock()
				// This should return a compatibility error.
				return SetupGenesisBlock(db, &customg)
			},
			wantHash:   customghash,
			wantConfig: customg.Config,
			wantErr: &params.ConfigCompatError{
				What:         "Homestead fork block",
				StoredConfig: big.NewInt(2),
				NewConfig:    big.NewInt(3),
				RewindTo:     1,
			},
		},
	}

	for _, test := range tests {
		db := ethdb.NewMemDatabase()
		config, hash, err := test.fn(db)
		// Check the return values.
		if !reflect.DeepEqual(err, test.wantErr) {
			spew := spew.ConfigState{DisablePointerAddresses: true, DisableCapacities: true}
			t.Errorf("%s: returned error %#v, want %#v", test.name, spew.NewFormatter(err), spew.NewFormatter(test.wantErr))
		}
		if !reflect.DeepEqual(config, test.wantConfig) {
			t.Errorf("%s:\nreturned %v\nwant     %v", test.name, config, test.wantConfig)
		}
		if hash != test.wantHash {
			t.Errorf("%s: returned hash %s, want %s", test.name, hash.Hex(), test.wantHash.Hex())
		} else if err == nil {
			// Check database content.
			stored := rawdb.ReadBlock(db, test.wantHash, 0)
			if stored.Hash() != test.wantHash {
				t.Errorf("%s: block in DB has hash %s, want %s", test.name, stored.Hash(), test.wantHash)
			}
		}
	}
}

func nonce2bytes(len int) []byte {
	a1 := len >> 16
	a2 := (len >> 8) & 0x00FF
	a3 := len & 0x0000FF
	buf := make([]byte, 3)
	buf[0] = uint8(a1)
	buf[1] = uint8(a2)
	buf[2] = uint8(a3)
	len2 := int(buf[0])*256*256 + int(buf[1])*256 + int(buf[2])
	if len != int(len2) {
		fmt.Println("len:", len, len2, buf[0], buf[1], buf[2])
		panic("[len2bytes] error code len")
	}
	return buf
}