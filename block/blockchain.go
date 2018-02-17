package block

//Blockchain is a chain of Block type, in this chain we can only add new blocks
//old blocks can not be altered
type Blockchain struct {
	blocks []*Block
}

//AddBlock add a new Block to the blockchain
func (bc *Blockchain) AddBlock(data StoredData) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//NewBlockchain returns a new Blockchain including the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{generateGenesisBlock()}}
}

func generateGenesisBlock() *Block {
	return NewBlock(&Bike{"InitialSerialNumber"}, "")
}
