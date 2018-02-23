package blockchain

import (
	"testing"
	"personal/GoBlockchain/persistance"
)

func TestNewBlockchain(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)
	blockchain := NewBlockchain(&fakePersistanceManager{})
	
	if blockchain == nil {
		t.Fatal("Failed to create new blockchain")
	}

	if blockchain.persistanceManager == nil {
		t.Error("Failed to create new blockchain")
	}
}

func TestAddBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)

	blockchain := NewBlockchain(&fakePersistanceManager{})

	err:= blockchain.AddBlock(&ConcreteData{"TestSerialNumber"})

	if err !=nil {
		t.Error("Failed to add new block")
	}
}


type fakePersistanceManager struct{}

func (m *fakePersistanceManager) SaveBlock(hash []byte, serializedBlock []byte, blockMetadata *persistance.BlockMetadata) error {

	return nil
}

func (m *fakePersistanceManager) RetrieveBlockByHash(hash []byte) ([]byte, error) {
	return []byte("Block"), nil
}

func (m *fakePersistanceManager) LastUsedHash() []byte {
	return []byte("testHash")
}

func (m *fakePersistanceManager) ClosePersistanceManager() {}
