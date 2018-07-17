/**
 * Package defining blocks in bungcoin
 * * * * */

package block

import (
	"time"
)

// A struct representing a block in the blockchain
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

// Function used to create a new block
func NewBlock(data string, prevBlockHash []byte) *Block {

	// Initialize the first 3 fields using args and current timestamp
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	// And return
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
