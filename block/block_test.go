package block

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)
	bike := Bike{"SN123545"}
	block, _ := NewBlock(&bike, []byte("SN123544"))
	blockData := block.Data.ToByteSlices()
	blockDataStr := string(blockData[:])
	if "SN123545" != blockDataStr {
		t.Errorf("Creation of new block failed new data: %v", blockDataStr)
	}
}
