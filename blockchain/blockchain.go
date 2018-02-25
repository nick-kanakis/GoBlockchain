package blockchain

import (
	"log"
	"personal/GoBlockchain/persistance"
)

const initialData = "Genesis block was generated"

//Blockchain is a chain of Block type, in this chain we can only add new blocks
//old blocks can not be altered
type Blockchain struct {
	persistanceManager persistance.Manager
}

//AddBlock add a new Block to the blockchain
func (bc *Blockchain) AddBlock(data []byte) error {
	lastBlock, err := getPreviousBlock(bc.persistanceManager)
	if err != nil {
		return err
	}
	newBlock, err := NewBlock(data, lastBlock)
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
func getPreviousBlock(pm persistance.Manager) (*Block, error) {
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
func NewBlockchain(address string, pm persistance.Manager) (*Blockchain, error) {
	genesisBlock := generateGenesisBlock(address)
	metadata := generateBlockMetadata(genesisBlock)
	genesisData, err := genesisBlock.Serialize()
	if err != nil {
		return nil, err
	}
	pm.SaveBlock(genesisBlock.Hash, genesisData, metadata)
	return &Blockchain{pm}, nil
}

func generateGenesisBlock(address string) *Block {
	block, err := NewGenesisBlock([]byte(initialData))

	if err != nil {
		log.Panicf("Could not incorporate genesis block into blockchain")
		return nil
	}

	return block
}

//NewBlockchainIterator returns a BlockchainIterator, the iteration is being done from newset to oldest block
func (bc *Blockchain) NewBlockchainIterator()*BlockchainIterator{
	return &BlockchainIterator{
		currentHash: bc.persistanceManager.LastUsedHash(),
		manager: bc.persistanceManager,
	}
}