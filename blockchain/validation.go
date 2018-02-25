package blockchain

import (
	"math/big"
	"crypto/sha256"
)

//ValidateBlock validates that the current block:
//1)has done work do be inserted in the blockchain
//2)previous hash is the same as the current value of previousBlockHash
//3)the height diff from previous block is 1
func ValidateBlock(last, current *Block) bool{
	
	if validateBlockHash(current) && validateHash(last, current) && validateHeight(last, current){
		return true
	}
	return false
}


func validateBlockHash(block *Block) bool {
	var bigHash big.Int
	pow := NewProofOfWork(block)
	headers := pow.newHeaders(block.Nonce)

	hash := sha256.Sum256(headers)
	bigHash.SetBytes(hash[:])

	if bigHash.Cmp(pow.target) == -1 {
		return true
	}
	return false
}

func validateHash(last, current *Block) bool{
	var lastHash, previousHash big.Int

	lastHash.SetBytes(current.PreviousBlockHash)
	previousHash.SetBytes(last.Hash)

	if lastHash.Cmp(&previousHash) != 0{
		return false
	}
	return true
}

func validateHeight(last, current *Block) bool{
	if last.Height + 1 == current.Height{
		return true
	}
	return false
}