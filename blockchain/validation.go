package blockchain

import (
	"math/big"
	"crypto/sha256"
)

//ValidateBlock validates that the current block:
//1)has done work do be inserted in the blockchain
//2)next olderst bock hash is the same as the current value of previousBlockHash
//3)the height diff from previous block is 1
func ValidateBlock(oldest, newest *Block) bool{	
	if validateBlockHash(newest) && validateHash(oldest, newest) && validateHeight(oldest, newest){
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

func validateHash(oldest, newest *Block) bool{
	var referenceHash, currentHash big.Int

	referenceHash.SetBytes(newest.PreviousBlockHash)
	currentHash.SetBytes(oldest.Hash)

	if currentHash.Cmp(&referenceHash) != 0{
		return false
	}
	return true
}

func validateHeight(oldest, newest *Block) bool{
	if oldest.Height + 1 == newest.Height{
		return true
	}
	return false
}