package block

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	bike := Bike{"SN123545"}
	block := NewBlock(&bike, "previousHash")
	
	if block.Data.ToString() != "SN123545"{
		t.Error("Creation of new block failed")
	}
}
