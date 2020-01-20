package main

import (
	"time"
)

// Block: store valuable information, contains some technical information,
// like its version, current timestamp and the hash of the previous block.

type Block struct {
	Timestamp     int64  // Block created time
	Data          []byte // Transactions
	PrevBlockHash []byte // The hash of the previous block
	Hash          []byte // The hash of the current block
	Nonce         int    // Require to verify the proof
}

// Hash algorithm
//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(header)
//	b.Hash = hash[:]
//}

// Block Constructor
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// To ensure every blockchain contains at least one block
// Dummy block as the head of the chain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
