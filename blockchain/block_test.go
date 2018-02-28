package blockchain

import (
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	data := []byte("testData")
	block, err := NewBlock(data, generateFakeBlock())
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}

	blockDataStr := string(block.Data)
	if "testData" != blockDataStr {
		t.Errorf("Creation of new block failed new data: %v", blockDataStr)
	}
}

func TestSerializeDeserializeBlock(t *testing.T) {
	data := []byte("testData")
	block, err := NewBlock(data, generateFakeBlock())
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}
	
	serializedBlock, err := block.Serialize()
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}

	deserializedBlock, err := DeserializeBlock(serializedBlock)
	if err != nil {
		t.Errorf("Could not deserialize block err msg: %v", err)
	}

	datastr := string(deserializedBlock.Data)
	if "testData" != datastr {
		t.Errorf("Serialization/Deserialization failed result was: %v", datastr)
	}
}

func TestAdjustedDifficulty(t *testing.T) {
	data := []byte("testData")
	block, err := NewBlock(data, generateFakeBlock())
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}

	for i := 0; i < difficultyAdjustmentInterval+1; i++ {
		block, err = NewBlock(data, block)
		if err != nil {
			t.Errorf("Could not create new block err msg: %v", err)
		}
	}

	if block.TargetBits == 8 {
		t.Error("Target bits should have been increased")
	}
}

func generateFakeBlock() *Block {
	return &Block{
		Timestamp:         time.Now().Unix(),
		Data:              []byte("fakeData"),
		PreviousBlockHash: []byte("fakePreviousBlockHash"),
		Hash:              []byte("fakeHash"),
		TargetBits:        8,
		Nonce:             1234,
		Height:            2,
	}
}
