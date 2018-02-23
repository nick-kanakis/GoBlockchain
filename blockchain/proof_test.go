package blockchain

import "testing"

func TestValidate(t *testing.T) {
	AdjustDifficulty(8)
	data := ConcreteData{"SN123545"}
	block, err := NewBlock(&data, []byte("SN123544"), 1)
	if err != nil {
		t.Errorf("Could not create new Block error msg: %v", err)
	}

	pow := NewProofOfWork(block)

	if !pow.Validate() {
		t.Error("Validation of block failed")
	}
}
