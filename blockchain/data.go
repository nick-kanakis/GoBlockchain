package blockchain

import (
	"bytes"
	"encoding/gob"
)

//StoredData is the interface that abstracts the concrete struct each
//block is storing
type StoredData interface {
	GetData() string
	Serialize() ([]byte, error)
}

//ConcreteData is the underling struct that the blockchain is build around
//at this point is only a string, in the future the ConcreteData will get
//more complicated
type ConcreteData struct {
	SerialNumber string
}

//GetData returns the underlying data
func (b ConcreteData) GetData() string {
	return b.SerialNumber
}

//Serialize implements the StoredData interface
func (b ConcreteData) Serialize() ([]byte, error) {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)

	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

//DeserializeData given a serialized CocreteData returns the deserialized struct
func DeserializeData(encodedData []byte) (*ConcreteData, error) {
	var data ConcreteData
	reader := bytes.NewReader(encodedData)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
