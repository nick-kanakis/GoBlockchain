package blockchain

import (
	"crypto/sha256"
	"errors"
	"log"
	"math"
	"math/big"
	"personal/GoBlockchain/utils"
)

//ErrFailedBlock is returned in case block fails to enter the blockchain
var ErrFailedBlock = errors.New("Failed to incorporate block into blockchain")

/*ProofOfWork defines the work that has to be done in order for
a new block to enter the blockchain. It has a Block pointer that refer to the
new block we want to add to the blockchain, also it has a target string.
target is used as the upper bound of the work that is needed to be done, more on that
in the Run() method.
*/
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

/*NewProofOfWork returns a new proof of work for that block,
here we specify the upper bound of the work. In the first 2 lines
we create a number where the (256-b.TargetBits) digit is 1 and before that
all digits are zero so: 00001XXXXXXXXX. This target will be used as a upper bound
for the work.

eg: If the work is to find a number (hash) that starts with at least
3 zeros, then target is: 001XXXXX...X and we need a number that is LESS than target.
*/
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.TargetBits))

	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

/*newHeaders is used to create a new new header based on:
- Hash of previous block
- Data of block
- Timestamp
- Target bits (set the difficulty)
- current nonce
The header will be used for next iteration of work
*/
func (pow *ProofOfWork) newHeaders(nonce uint) []byte {

	return utils.ConcatByteSlices(
		pow.block.PreviousBlockHash,
		pow.block.Data.ToByteSlices(),
		utils.UintToByteSlice(uint64(pow.block.Timestamp)),
		utils.UintToByteSlice(uint64(pow.block.TargetBits)),
		utils.UintToByteSlice(uint64(nonce)),
	)
}

/*DoWork is the function that does the work needed to add the block into the blockchain,
We start with a nonce = 0 and we prepare a header to be used in the new hash function.
In order for this function to be correct it must be less than the target.
Each time the hash fails to meet the requirement stated previous the nonce is increased and
the hash is tried again.
*/
func (pow *ProofOfWork) DoWork() (uint, []byte, error) {
	var currentHash big.Int
	var maxNonce uint = math.MaxInt64
	var hash [32]byte
	var nonce uint

	for nonce = 0; nonce < maxNonce; nonce++ {
		headers := pow.newHeaders(nonce)
		hash = sha256.Sum256(headers)

		currentHash.SetBytes(hash[:])

		//currentHash < target
		if currentHash.Cmp(pow.target) == -1 {
			return nonce, hash[:], nil
		}
	}

	log.Panicf("Nonce exceed math.MaxInt64, block %v did not entered the blockchain ", pow.block.Data)
	return 0, []byte{}, ErrFailedBlock
}

//Validate is used to validate that a block has done the necessary work in order
//to be part of the block
func (pow *ProofOfWork) Validate() bool {
	var bigHash big.Int
	headers := pow.newHeaders(pow.block.Nonce)
	hash := sha256.Sum256(headers)
	bigHash.SetBytes(hash[:])

	if bigHash.Cmp(pow.target) == -1 {
		return true
	}
	return false
}
