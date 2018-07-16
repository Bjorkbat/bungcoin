/**
 * Package for describing functionality of the blockchain itself (not the
 * blocks though, see the block package for that)
 * * * * */

package blockchain

import (
	"github.com/Bjorkbat/bungcoin/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

// Function to add blocks to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}
