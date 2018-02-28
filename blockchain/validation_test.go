package blockchain

import (
	"testing"
)

func TestValidateBlock(t *testing.T) {
	oldestBlock, err := NewBlock([]byte("testData"), generateFakeBlock())
	newestBlock, err := NewBlock([]byte("testData2"), oldestBlock)
	if err != nil {
		t.Errorf("Could not create new Block error msg: %v", err)
	}

	if !ValidateBlock(oldestBlock, newestBlock) {
		t.Error("Validation of block failed")
	}
}
