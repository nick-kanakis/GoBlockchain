package blockchain

import "testing"

func TestValidate(t *testing.T) {
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, _ := NewBlock(&data, []byte("SN123544"), 1)
	pow := NewProofOfWork(block)

	if !pow.Validate() {
		t.Error("Validation of block failed")
	}
}
