package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// The difficulty mined this block at the block header
// To simplify, there is no target adjusting algorithm
// Target is an upper bound of a range
const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// Constructor
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

// Merge block fields with the target and nonce
// nonce: A counter from the Hashcash
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//Implementation of Algorithm
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int // Integer representation of hash
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce { // Avoid a possible overflow of nonce
		data := pow.prepareData(nonce)// Prepare data
		hash = sha256.Sum256(data)// hash with SHA256
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])// convert to big int

		if hashInt.Cmp(pow.target) == -1 {// compare with target
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// Valide the POW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
