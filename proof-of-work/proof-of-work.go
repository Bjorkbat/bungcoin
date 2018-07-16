/**
 * Package defining our proof of work mechanism.
 * * * * */

import (
	"math/big"
)

// Globally defined difficulty.  No need to adjust the difficulty
// algorithmically for now
// FYI, the target is arbitrary, I might even change it later
const TARGET_BITS = 24

type ProofOfWork struct {
	block *block.Block
	target *big.Int
}

func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))

	pow := &ProofOfWork{b, target}

	return pow
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
