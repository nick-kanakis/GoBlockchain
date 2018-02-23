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
	if err !=nil{
		log.Println("Could not fetch block from db")
		return err
	}
	newBlock, err := NewBlock(data, lastBlock.Hash, lastBlock.Height)
	if err != nil {
		log.Printf("Could not incorporate block %v into blockchain\n", newBlock.Data)
		return err
	}
	newBlockMetadata:= generateBlockMetadata(newBlock)
	bc.persistanceManager.SaveBlock(newBlock.Hash, newBlock.Serialize(), newBlockMetadata)
	
	return nil
}
//Retrieve previous Block
func getPreviousHashHeight(pm persistance.Manager) (*Block, error){
	lastHash := pm.LastUsedHash()
	serializedBlock, err := pm.RetrieveBlockByHash(lastHash)
	
	if err != nil{
		return nil, err
	}
	block := DeserializeBlock(serializedBlock)

	return block, nil
}

func generateBlockMetadata(block *Block) *persistance.BlockMetadata{
	return &persistance.BlockMetadata{block.Height, ""}
}

//NewBlockchain returns a new Blockchain including the genesis block
func NewBlockchain(pm persistance.Manager) *Blockchain {
	genesisBlock:=generateGenesisBlock()
	metadata:= generateBlockMetadata(genesisBlock)
	pm.SaveBlock(genesisBlock.Hash, genesisBlock.Serialize(), metadata)
	return &Blockchain{pm}
}

func generateGenesisBlock() *Block {
	block, err := NewBlock(&ConcreteData{"InitialSerialNumber"}, []byte{}, -1)

	if err != nil {
		log.Panicf("Could not incorporate genesis block into blockchain")
		return nil
	}

	return block
}
