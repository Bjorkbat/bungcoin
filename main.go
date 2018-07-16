/**
 * The results of a simple blockchain tutorial I followed to finally get the
 * hands-on technical understanding I've been wanting.
 * Bung is many things, including a technical term for a pig's anus.  This
 * hypthothetical coin is used to track one's ownership of pig anus assets.
 */

package main

import (
	"fmt"

	"github.com/Bjorkbat/bungcoin/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 Bungcoin to John")
	bc.AddBlock("Send 2 more Bungcoin to John")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
