package blockchain

import (
	"personal/GoBlockchain/persistance"
)

//Iterator helps the iteration of the blockchain, the blocks are iterated
//from the most newly added to the oldest. This means that the last block
//will be the genesis block. 
type Iterator struct {
	currentHash []byte
	manager     persistance.Manager
}

//Next returns next block in blockchain.
func (iter *Iterator) Next() (*Block, error) {
	encodedBlock, err := iter.manager.RetrieveBlockByHash(iter.currentHash)
	if err != nil {
		return nil, err
	}
	block, err := DeserializeBlock(encodedBlock)
	if err != nil {
		return nil, err
	}
	iter.currentHash = block.PreviousBlockHash

	return block, nil
}
