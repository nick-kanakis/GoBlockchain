package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

//StoredData is the interface that abstracts the concrete struct each
//block is storing
type StoredData interface {
	GetData() string
	Serialize() []byte
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
func (b ConcreteData) Serialize() []byte{
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)

	err := encoder.Encode(b)
	if err != nil {
		log.Panicf("Could not serialize data error msg: %v", err)
	}
	return  buff.Bytes()
}

//DeserializeData given a serialized CocreteData returns the deserialized struct
func DeserializeData(encodedData []byte) *ConcreteData{
	var data ConcreteData
	reader := bytes.NewReader(encodedData)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&data)

	if err!=nil{
		log.Panicf("Could not deserialize data error msg: %v", err)
	}

	return &data
}
