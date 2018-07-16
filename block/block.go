/**
 * Package defining blocks in bungcoin
 * * * * */

package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// A struct representing a block in the blockchain
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

// Function used to create a new block
func NewBlock(data string, prevBlockHash []byte) *Block {

	// Initialize the first 3 fields using args and current timestamp
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}

	// Set the hash on our block
	block.SetHash()

	// And return
	return block
}

// Function used to set the hash on a block
func (b *Block) SetHash() {
	// Take the timestamp and convert it into a string.  From there we can easily
	// turn it into a byte array
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// What's this gnarly thing?  Well, creating an array of byte arrays
	// (technically slices) and joining these arrays of arrays using bytes.Join
	// That second arg is typically reserved for seperators.  We're saying here
	// that we won't want any seperators here by passing in an empty byte array
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	// And he we create our hash
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
