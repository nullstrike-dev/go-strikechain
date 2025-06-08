package vault

import "crypto/ecdsa"

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	return &Wallet{}
}
