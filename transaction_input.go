package main

import "bytes"

type TXInput struct {
	Txid      []byte // Transaction ID
	Vout      int    // Index of output in the transaction
	Signature []byte
	PubKey    []byte
}

// Check that an input uses a specific key to unlock an output
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
