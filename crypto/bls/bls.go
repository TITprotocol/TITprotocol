// Package bls implements a go-wrapper around a library implementing the
// the BLS12-381 curve and signature scheme. This package exposes a public API for
// verifying and aggregating BLS signatures used by Ethereum 2.0.
package bls

import (
	"fmt"
	"io"

	g1 "github.com/phoreproject/bls/g1pubs"
	"github.com/TITprotocol/go-TITprotocol/common/bytesutil"
)

// Signature used in the BLS signature scheme.
type Signature struct {
	val *g1.Signature
}

// SecretKey used in the BLS signature scheme.
type SecretKey struct {
	val *g1.SecretKey
}

// PublicKey used in the BLS signature scheme.
type PublicKey struct {
	val *g1.PublicKey
}

// RandKey creates a new private key using a random method provided as an io.Reader.
func RandKey(r io.Reader) (*SecretKey, error) {
	k, err := g1.RandKey(r)
	if err != nil {
		return nil, fmt.Errorf("could not initialize secret key: %v", err)
	}
	return &SecretKey{val: k}, nil
}

// SecretKeyFromBytes creates a BLS private key from a byte slice.
func SecretKeyFromBytes(priv []byte) (*SecretKey, error) {
	if len(priv) != 32 {
		return nil, fmt.Errorf("expected byte slice of length 32, received: %d", len(priv))
	}
	k := bytesutil.ToBytes32(priv)
	return &SecretKey{val: g1.DeserializeSecretKey(k)}, nil
}

// PublicKeyFromBytes creates a BLS public key from a byte slice.
func PublicKeyFromBytes(pub []byte) (*PublicKey, error) {
	b := bytesutil.ToBytes48(pub)
	k, err := g1.DeserializePublicKey(b)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal bytes into public key: %v", err)
	}
	return &PublicKey{val: k}, nil
}

// SignatureFromBytes creates a BLS signature from a byte slice.
func SignatureFromBytes(sig []byte) (*Signature, error) {
	b := bytesutil.ToBytes96(sig)
	s, err := g1.DeserializeSignature(b)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal bytes into signature: %v", err)
	}
	return &Signature{val: s}, nil
}

// PublicKey obtains the public key corresponding to the BLS secret key.
func (s *SecretKey) PublicKey() *PublicKey {
	return &PublicKey{val: g1.PrivToPub(s.val)}
}

// Sign a message using a secret key - in a beacon/validator client,
func (s *SecretKey) Sign(msg []byte, domain uint64) *Signature {
	sig := g1.SignWithDomain(bytesutil.ToBytes32(msg), s.val, domain)
	return &Signature{val: sig}
}

// Marshal a secret key into a byte slice.
func (s *SecretKey) Marshal() []byte {
	k := s.val.Serialize()
	return k[:]
}

// Marshal a public key into a byte slice.
func (p *PublicKey) Marshal() []byte {
	k := p.val.Serialize()
	return k[:]
}

// Aggregate two public keys.
func (p *PublicKey) Aggregate(p2 *PublicKey) *PublicKey {
	p1 := p.val
	p1.Aggregate(p2.val)
	return &PublicKey{val: p1}
}

// Verify a bls signature given a public key, a message, and a domain.
func (s *Signature) Verify(msg []byte, pub *PublicKey, domain uint64) bool {
	return g1.VerifyWithDomain(bytesutil.ToBytes32(msg), pub.val, s.val, domain)
}

// VerifyAggregate verifies each public key against a message.
// This is vulnerable to rogue public-key attack. Each user must
// provide a proof-of-knowledge of the public key.
func (s *Signature) VerifyAggregate(pubKeys []*PublicKey, msg []byte, domain uint64) bool {
	if len(pubKeys) == 0 {
		return false // Otherwise panic in VerifyAggregateCommonWithDomain.
	}
	var keys []*g1.PublicKey
	for _, v := range pubKeys {
		keys = append(keys, v.val)
	}
	return s.val.VerifyAggregateCommonWithDomain(keys, bytesutil.ToBytes32(msg), domain)
}

// Marshal a signature into a byte slice.
func (s *Signature) Marshal() []byte {
	k := s.val.Serialize()
	return k[:]
}

// AggregateSignatures converts a list of signatures into a single, aggregated sig.
func AggregateSignatures(sigs []*Signature) *Signature {
	var ss []*g1.Signature
	for _, v := range sigs {
		ss = append(ss, v.val)
	}
	return &Signature{val: g1.AggregateSignatures(ss)}
}

// Domain returns the bls domain given by the domain type and the operation 4 byte fork version.
//
// Spec pseudocode definition:
//  def get_domain(state: MasternodeState, domain_type: DomainType, message_epoch: Epoch=None) -> Domain:
//    """
//    Return the signature domain (fork version concatenated with domain type) of a message.
//    """
//    epoch = get_current_epoch(state) if message_epoch is None else message_epoch
//    fork_version = state.fork.previous_version if epoch < state.fork.epoch else state.fork.current_version
//    return compute_domain(domain_type, fork_version)
func Domain(domainType []byte, forkVersion []byte) uint64 {
	b := []byte{}
	b = append(b, domainType[:4]...)
	b = append(b, forkVersion[:4]...)
	return bytesutil.FromBytes8(b)
}
