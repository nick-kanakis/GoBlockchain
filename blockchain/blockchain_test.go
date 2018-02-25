package blockchain

import (
	"personal/GoBlockchain/persistance"
	"testing"
	"time"
)

func TestNewBlockchain(t *testing.T) {
	blockchain, err := NewBlockchain("testAddress",&fakePersistanceManager{})
	if err != nil {
		t.Errorf("Could not create new Blockchain error msg: %v", err)
	}
	if blockchain == nil {
		t.Fatal("Failed to create new blockchain")
	}

	if blockchain.persistanceManager == nil {
		t.Error("Failed to create new blockchain")
	}
}

func TestAddBlock(t *testing.T) {
	blockchain, err := NewBlockchain("testAddress",&fakePersistanceManager{})
	if err != nil {
		t.Errorf("Could not create new Blockchain error msg: %v", err)
	}

	err = blockchain.AddBlock([]byte("TestSerialNumber"))
	if err != nil {
		t.Errorf("Failed to add new block error msg: %v", err)
	}
}

func TestNewBlockchainIterator(t *testing.T) {
	blockchain, err := NewBlockchain("testAddress",&fakePersistanceManager{})
	if err != nil {
		t.Errorf("Could not create new Blockchain error msg: %v", err)
	}
	iter := blockchain.NewIterator()
	previousBlock, err := iter.Next()
	if err != nil {
		t.Errorf("Could not iterate previous block error msg: %v", err)
	}

	if previousBlock.Data == nil {
		t.Errorf("Iteration failed previous block has corrupted data error msg: %v", err)
	}
}

type fakePersistanceManager struct{}

func (m *fakePersistanceManager) SaveBlock(hash []byte, serializedBlock []byte, blockMetadata *persistance.BlockMetadata) error {
	return nil
}

func (m *fakePersistanceManager) RetrieveBlockByHash(hash []byte) ([]byte, error) {
	block := Block{
		Timestamp:         time.Now().Unix(),
		Data:              []byte("testData"),
		PreviousBlockHash: []byte("1234"),
		Hash:              []byte{},
		TargetBits:        8,
		Nonce:             0,
		Height:            1,
	}
	serializedBlock, _ := block.Serialize()
	return serializedBlock, nil
}

func (m *fakePersistanceManager) LastUsedHash() []byte {
	return []byte("testHash")
}

func (m *fakePersistanceManager) ClosePersistanceManager() error {
	return nil
}
