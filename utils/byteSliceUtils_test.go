package utils

import (
	"encoding/binary"
	"testing"
)

func TestConcatByteSlicesMultipleInputs(t *testing.T) {
	var slice1 = []byte{0XFF, 0XFE}
	var slice2 = []byte{0XFD, 0XFC}
	var slice3 = []byte{0XFA}

	concatSlices := ConcatByteSlices(slice1, slice2, slice3)
	expectedResult := []byte{0XFF, 0XFE, 0XFD, 0XFC, 0XFA}

	for i := range concatSlices {
		if concatSlices[i] != expectedResult[i] {
			t.Error("Concatenation of strings failed")
		}
	}
}

func TestConcatByteSlicesSingleInputs(t *testing.T) {
	var slice1 = []byte{0XFF, 0XFE}

	concatSlices := ConcatByteSlices(slice1)
	expectedResult := []byte{0XFF, 0XFE}

	for i := range concatSlices {
		if concatSlices[i] != expectedResult[i] {
			t.Error("Concatenation of strings failed")
		}
	}
}

func TestEpochToByteSlice(t *testing.T) {
	var timestamp int64 = 1518973818
	timestampToBytes := EpochToByteSlice(timestamp)

	if "1518973818" != string(timestampToBytes[:]) {
		t.Fatalf("Epoch time convertion fail result has: %v", timestampToBytes)
	}
}

func TestUintToByteSlice(t *testing.T) {
	slice := UintToByteSlice(1234)

	result := binary.BigEndian.Uint64(slice)
	if result != 1234 {
		t.Fatalf("Uint to byte slice convertion fail result has: %v", slice)
	}
}
