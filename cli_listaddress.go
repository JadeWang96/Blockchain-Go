package main

import (
	"fmt"
)

func (cli *CLI) listAddresses() {
	wallets := NewWallet()

	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}