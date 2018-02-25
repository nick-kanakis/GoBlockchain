package blockchain

import (
	"testing"
)
func TestValidateBlock(t *testing.T) {
	previousBlock, err := NewBlock([]byte("testData"), generateFakeBlock())
	currentBlock, err := NewBlock([]byte("testData2"),  previousBlock)
	if err != nil {
		t.Errorf("Could not create new Block error msg: %v", err)
	}

	if !ValidateBlock(previousBlock, currentBlock) {
		t.Error("Validation of block failed")
	}
}