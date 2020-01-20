package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Println("Prev. hash: %x", block.PrevBlockHash)
		fmt.Println("Data: %s", block.Data)
		fmt.Println("Hash: %s", block.Hash)
	}
}
