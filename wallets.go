package main

// Keep a collection of wallets, save them to a file, and load them from it
type Wallets struct {
	Wallets map[string]*Wallet
}

