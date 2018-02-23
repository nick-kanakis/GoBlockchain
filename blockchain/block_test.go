package blockchain

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	//adjust diff to be quite low (only 8 bits)
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, err := NewBlock(&data, []byte("SN123544"), 1)
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}

	blockDataStr := block.Data.GetData()
	if "SN123545" != blockDataStr {
		t.Errorf("Creation of new block failed new data: %v", blockDataStr)
	}
}

func TestSerializeDeserializeBlock(t *testing.T) {
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, _ := NewBlock(&data, []byte("SN123544"), 1)
	serializedBlock, err := block.Serialize()
	if err != nil {
		t.Errorf("Could not create new block err msg: %v", err)
	}

	deserializedBlock, err := DeserializeBlock(serializedBlock)
	if err != nil {
		t.Errorf("Could not deserialize block err msg: %v", err)
	}

	datastr := deserializedBlock.Data.GetData()
	if "SN123545" != datastr {
		t.Errorf("Serialization/Deserialization failed result was: %v", datastr)
	}
}
