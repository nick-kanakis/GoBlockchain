package blockchain

import (
	"testing"
)

func TestSerializeDeserializeData(t *testing.T){

	data := &ConcreteData{"SR1234"}
	serializedData := data.Serialize()

	deserializedData :=DeserializeData(serializedData)
	deserializedDataStr := deserializedData.GetData()
	if "SR1234" != deserializedDataStr{
		t.Errorf("Serialization/Deserialization failed result was: %v", deserializedDataStr)
	}
}