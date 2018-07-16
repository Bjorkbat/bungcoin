/**
 * The results of a simple blockchain tutorial I followed to finally get the
 * hands-on technical understanding I've been wanting.
 * Bung is many things, including a technical term for a pig's anus.  This
 * hypthothetical coin is used to track one's ownership of pig anus assets.
 */

package bungcoin

// A struct representing a block in the blockchain
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}
