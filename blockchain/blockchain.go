package blockchain

import (
	"log"
	"personal/GoBlockchain/persistance"
)

//Blockchain is a chain of Block type, in this chain we can only add new blocks
//old blocks can not be altered
type Blockchain struct {
	persistanceManager persistance.Manager
}

//AddBlock add a new Block to the blockchain
func (bc *Blockchain) AddBlock(data StoredData) error {
	lastBlock, err := getPreviousHashHeight(bc.persistanceManager)
	if err != nil {
		return err
	}
	newBlock, err := NewBlock(data, lastBlock.Hash, lastBlock.Height)
	if err != nil {
		return err
	}
	newBlockMetadata := generateBlockMetadata(newBlock)
	serializedData, err := newBlock.Serialize()
	if err != nil {
		return err
	}
	bc.persistanceManager.SaveBlock(newBlock.Hash, serializedData, newBlockMetadata)

	return nil
}

//Retrieve previous Block
func getPreviousHashHeight(pm persistance.Manager) (*Block, error) {
	lastHash := pm.LastUsedHash()
	serializedBlock, err := pm.RetrieveBlockByHash(lastHash)

	if err != nil {
		return nil, err
	}
	block, err := DeserializeBlock(serializedBlock)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func generateBlockMetadata(block *Block) *persistance.BlockMetadata {
	return &persistance.BlockMetadata{block.Height, ""}
}

//NewBlockchain returns a new Blockchain including the genesis block
func NewBlockchain(pm persistance.Manager) (*Blockchain, error) {
	genesisBlock := generateGenesisBlock()
	metadata := generateBlockMetadata(genesisBlock)
	genesisData, err := genesisBlock.Serialize()
	if err != nil {
		return nil, err
	}
	pm.SaveBlock(genesisBlock.Hash, genesisData, metadata)
	return &Blockchain{pm}, nil
}

func generateGenesisBlock() *Block {
	block, err := NewBlock(&ConcreteData{"InitialSerialNumber"}, []byte{}, -1)

	if err != nil {
		log.Panicf("Could not incorporate genesis block into blockchain")
		return nil
	}

	return block
}
