package blockchain

import (
	"testing"
)

func TestSerializeDeserializeData(t *testing.T) {

	data := &ConcreteData{"SR1234"}
	serializedData, err := data.Serialize()
	if err != nil {
		t.Errorf("Could not serliaze data error msg: %v", err)
	}
	deserializedData, err := DeserializeData(serializedData)
	if err != nil {
		t.Errorf("Could not deserialize data error msg: %v", err)
	}

	deserializedDataStr := deserializedData.GetData()
	if "SR1234" != deserializedDataStr {
		t.Errorf("Serialization/Deserialization failed result was: %v", deserializedDataStr)
	}
}
