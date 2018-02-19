package utils

import (
	"encoding/binary"
)

//ConcatByteSlices concatenate multiple byte slices into one
func ConcatByteSlices(byteSlices ...[]byte) []byte {
	var concatenated []byte

	for _, slice := range byteSlices {
		concatenated = append(concatenated, slice...)
	}
	return concatenated
}

//UintToByteSlice converts int64 to byte slices
func UintToByteSlice(num uint64) []byte {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, num)
	return buffer
}
