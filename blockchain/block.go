package blockchain

import (
	"bytes"
	"encoding/gob"
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
	Data      []byte
	PreviousBlockHash []byte
	Hash              []byte
	TargetBits        uint
	Nonce             uint
	Height            int
}

//NewBlock creates a new block based on the previous block.
func NewBlock(data []byte, previousBlock *Block) (*Block, error) {
	block := &Block{
		Timestamp:         time.Now().Unix(),
		Data:      data,
		PreviousBlockHash: previousBlock.Hash,
		Hash:              []byte{},
		TargetBits:        difficulty,
		Nonce:             0,
		Height:            previousBlock.Height + 1,
	}
	pow := NewProofOfWork(block)
	nonce, hash, err := pow.DoWork()

	if err != nil {
		return nil, err
	}

	block.Hash = hash
	block.Nonce = nonce
	return block, nil
}

//NewGenesisBlock generates the first block in the blockchain
func NewGenesisBlock(data []byte) (*Block, error) {
	block := &Block{
		Timestamp:         time.Now().Unix(),
		Data:      		   data,
		PreviousBlockHash: []byte{},
		Hash:              []byte{},
		TargetBits:        difficulty,
		Nonce:             0,
		Height:            0,
	}
	pow := NewProofOfWork(block)
	nonce, hash, err := pow.DoWork()

	if err != nil {
		return nil, err
	}

	block.Hash = hash
	block.Nonce = nonce
	return block, nil
}

//DeserializeBlock deseriaize a block from a byte slice
func DeserializeBlock(encodedBlock []byte) (*Block, error) {
	var block Block
	reader := bytes.NewReader(encodedBlock)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

//Serialize serialize the block to a slice
func (b *Block) Serialize() ([]byte, error) {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(b)

	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

//AdjustDifficulty set the difficulty of creating a new block
func AdjustDifficulty(diff uint) {
	difficulty = diff
}
