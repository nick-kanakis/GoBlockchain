package blockchain

import (
	"log"
	"personal/GoBlockchain/persistance"
	"fmt"
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
func NewBlockchain(pm persistance.Manager) (*Blockchain, error) {
	lastUsedHash:= pm.LastUsedHash()
	if len(lastUsedHash) == 0{
		genesisBlock := generateGenesisBlock()
		metadata := generateBlockMetadata(genesisBlock)
		genesisData, err := genesisBlock.Serialize()
		if err != nil {
			return nil, err
		}
		pm.SaveBlock(genesisBlock.Hash, genesisData, metadata)
	}
	return &Blockchain{pm}, nil
}

func generateGenesisBlock() *Block {
	block, err := NewGenesisBlock([]byte(initialData))

	if err != nil {
		log.Panicf("Could not incorporate genesis block into blockchain")
		return nil
	}

	return block
}

//NewIterator returns a blockchain Iterator,
//the iteration is being done from newset to oldest block
func (bc *Blockchain) NewIterator()*Iterator{
	return &Iterator{
		currentHash: bc.persistanceManager.LastUsedHash(),
		manager: bc.persistanceManager,
	}
}

func (bc *Blockchain) PrintChain() error{
	iter := bc.NewIterator()
	block, err := iter.Next()
	
	for block != nil{
		if err != nil{
			return err
		}
		fmt.Println(block)
		block, err = iter.Next()
	}
	return nil
}

var validate = ValidateBlock
//ValidateChain validates that each block in the chain
//is a valid block and is in a valid position
func (bc *Blockchain) ValidateChain() bool{
	iter := bc.NewIterator()
	newest, err := iter.Next()
	if err != nil{
		return false
	}
	oldest, err := iter.Next()
	
	for oldest != nil{
		if err != nil{
			return false
		}

		if !validate(oldest, newest){
			return false
		}
		newest = oldest
		oldest, err = iter.Next()
	}
	return true
}