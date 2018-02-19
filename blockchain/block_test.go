package blockchain

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, _ := NewBlock(&data, []byte("SN123544"))
	blockDataStr := block.Data.GetData()
	if "SN123545" != blockDataStr {
		t.Errorf("Creation of new block failed new data: %v", blockDataStr)
	}
}

func TestSerializeDeserializeBlock(t *testing.T) {
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, _ := NewBlock(&data, []byte("SN123544"))
	serializedBlock := block.Serialize()
	deserializedBlock := DeserializeBlock(serializedBlock)
	datastr := deserializedBlock.Data.GetData()

	if "SN123545" != datastr {
		t.Errorf("Serialization/Deserialization failed result was: %v", datastr)
	}
}
