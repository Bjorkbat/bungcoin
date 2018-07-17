/**
 * Part of the block package defining how our proof of work, well, works.
 * * * * */

package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// Globally defined difficulty.  No need to adjust the difficulty
// algorithmically for now
// FYI, the target is arbitrary, I might even change it later
const TARGET_BITS = 24

const MAX_NONCE = math.MaxInt64

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(TARGET_BITS)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))

	pow := &ProofOfWork{b, target}

	return pow
}


func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash[32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < MAX_NONCE {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break;
		} else {
			nonce ++
		}

		fmt.Print("\n\n")
	}

	return nonce, hash[:]

}

func IntToHex(i int64) []byte {
	return []byte(strconv.FormatInt(i, 16))
}

func (pow *ProofOfWork) Validate() bool {

	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid

}
