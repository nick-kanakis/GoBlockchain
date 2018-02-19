package blockchain

import "testing"

func TestNewBlockchain(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)
	blockchain := NewBlockchain()
	if blockchain == nil {
		t.Fatal("Failed to create new blockchain")
	}

	if len(blockchain.blocks) != 1 {
		t.Error("Failed to add genesis block")
	}
}

func TestAddBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)

	blockchain := NewBlockchain()

	blockchain.AddBlock(&ConcreteData{"TestSerialNumber"})

	if len(blockchain.blocks) != 2 {
		t.Error("Failed to add new block")
	}
	lastBlockDataStr := blockchain.blocks[len(blockchain.blocks)-1].Data.GetData()
	if lastBlockDataStr != "TestSerialNumber" {
		t.Error("New block was not added correctly")
	}
}
