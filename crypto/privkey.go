// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crypto

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

var (
	secp256k1Curve = btcec.S256()
)

// PrivateKey is a btcec.PrivateKey wrapper
type PrivateKey btcec.PrivateKey

// KeyPairFromBytes returns a private and public key pair from private key passed as a byte slice privKeyBytes
func KeyPairFromBytes(privKeyBytes []byte) (*PrivateKey, *PublicKey, error) {
	if len(privKeyBytes) != btcec.PrivKeyBytesLen {
		return nil, nil, fmt.Errorf("Private key must be be exactly %d bytes (%d)", btcec.PrivKeyBytesLen, len(privKeyBytes))
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(secp256k1Curve, privKeyBytes)
	return (*PrivateKey)(privKey), (*PublicKey)(pubKey), nil
}

// NewKeyPair returns a new private and public key pair
func NewKeyPair() (*PrivateKey, *PublicKey, error) {
	btcecPrivKey, err := btcec.NewPrivateKey(secp256k1Curve)
	if err != nil {
		return nil, nil, err
	}
	privKey := (*PrivateKey)(btcecPrivKey)
	return privKey, privKey.PubKey(), nil
}

// Serialize convert private key into byte array
func (p *PrivateKey) Serialize() []byte {
	return ((*btcec.PrivateKey)(p)).Serialize()
}

// PubKey returns the PublicKey corresponding to this private key.
func (p *PrivateKey) PubKey() *PublicKey {
	return (*PublicKey)((*btcec.PrivateKey)(p).PubKey())
}

// Erase destroys the private info of private key, and leaves only public info
func (p *PrivateKey) Erase() {
	bits := p.D.Bits()
	for i := 0; i < len(bits); i++ {
		bits[i] = 0
	}
}
