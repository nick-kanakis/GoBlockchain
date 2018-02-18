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

	blockchain.AddBlock(&Bike{"TestSerialNumber"})

	if len(blockchain.blocks) != 2 {
		t.Error("Failed to add new block")
	}
	lastBlockData := blockchain.blocks[len(blockchain.blocks)-1].Data.ToByteSlices()
	lastBlockDataStr := string(lastBlockData[:])
	if lastBlockDataStr != "TestSerialNumber" {
		t.Error("New block was not added correctly")
	}
}
