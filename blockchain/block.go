package blockchain

import (
	"log"
	"time"
)

//This need to be dynamically adjusted
//Currently it has 24 bits or 3 bytes
var difficulty uint = 24

//Block represents the block of the chain it is composed by the following:
//Timestamp: the time (in Unix time) that the block was created
//Data: the source data that the block is storing
//PreviousBlockHash: Hash value of the previous block
//Hash: Hash value of the current block
//TargetBits: define the difficulty of work that need to be done in order a block to
//enter the blockchain
//nonce: number of tries it took for this block to enter the blockchain
type Block struct {
	Timestamp         int64
	Data              StoredData
	PreviousBlockHash []byte
	Hash              []byte
	TargetBits        uint
	Nonce             uint
}

//NewBlock creates a new block based on the previous block.
func NewBlock(data StoredData, prevBlockHash []byte) (*Block, error) {
	block := &Block{
		Timestamp:         time.Now().Unix(),
		Data:              data,
		PreviousBlockHash: prevBlockHash,
		Hash:              []byte{},
		TargetBits:        difficulty,
		Nonce:             0,
	}
	pow := NewProofOfWork(block)
	nonce, hash, err := pow.DoWork()

	if err != nil {
		log.Panicf("Could not incorporate block %v into blockchain", block.Data)
		return nil, err
	}

	block.Hash = hash
	block.Nonce = nonce
	return block, nil
}

//AdjustDifficulty set the difficulty of creating a new block
func AdjustDifficulty(diff uint) {
	difficulty = diff
}
