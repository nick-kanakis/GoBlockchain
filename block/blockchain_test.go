package block

import "testing"

func TestNewBlockchain(t *testing.T){
	blockchain := NewBlockchain()
	
	if blockchain == nil {
		t.Fatal("Failed to create new blockchain")
	}

	if len(blockchain.blocks) != 1 {
		t.Error("Failed to add genesis block")
	}
}

func TestAddBlock(t *testing.T){
	blockchain := NewBlockchain()
	
	blockchain.AddBlock(&Bike{"TestSerialNumber"})
	
	if len(blockchain.blocks) != 2 {
		t.Error("Failed to add new block")
	}

	if blockchain.blocks[len(blockchain.blocks) -1].Data.ToString() != "TestSerialNumber"{
		t.Error("New block was not added correctly")
	}
}