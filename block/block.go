package block

import (
	"crypto/sha256"
	"encoding/hex"
	"personal/GoBlockchain/utils"
	"time"
)

//Block represents the block of the chain it is composed by the following:
//Timestamp: the time (in Unix time) that the block was created
//Data: the source data that the block is storing
//PreviousBlockHash: Hash value of the previous block
//Hash: Hash value of the current block
type Block struct {
	Timestamp         int64
	Data              StoredData
	PreviousBlockHash string
	Hash              string
}

func (b *Block) calculateHash() {
	headers := utils.ConcatStrings(string(b.Timestamp), string(b.PreviousBlockHash), b.Data.ToString())
	hash := sha256.Sum256([]byte(headers))
	b.Hash = hex.EncodeToString(hash[:])
}

//NewBlock creates a new block based on the previous block.
func NewBlock(data StoredData, prevBlockHash string) *Block {
	block := &Block{
		Timestamp:         time.Now().Unix(),
		Data:              data,
		PreviousBlockHash: prevBlockHash,
		Hash:              "",
	}
	block.calculateHash()
	return block
}
