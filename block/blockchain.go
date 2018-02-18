package block

import (
	"log"
)

//Blockchain is a chain of Block type, in this chain we can only add new blocks
//old blocks can not be altered
type Blockchain struct {
	blocks []*Block
}

//AddBlock add a new Block to the blockchain
func (bc *Blockchain) AddBlock(data StoredData) error {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock, err := NewBlock(data, lastBlock.Hash)
	if err != nil {
		log.Panicf("Could not incorporate block %v into blockchain", newBlock.Data)
		return err
	}
	bc.blocks = append(bc.blocks, newBlock)
	return nil
}

//NewBlockchain returns a new Blockchain including the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{generateGenesisBlock()}}
}

func generateGenesisBlock() *Block {
	block, err := NewBlock(&Bike{"InitialSerialNumber"}, []byte{})

	if err != nil {
		log.Panicf("Could not incorporate genesis block into blockchain")
		return nil
	}

	return block
}
