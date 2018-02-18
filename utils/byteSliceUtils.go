package utils

import (
	"encoding/binary"
	"strconv"
)

//ConcatByteSlices concatenate multiple byte slices into one
func ConcatByteSlices(byteSlices ...[]byte) []byte {
	var concatenated []byte

	for _, slice := range byteSlices {
		concatenated = append(concatenated, slice...)
	}
	return concatenated
}

//EpochToByteSlice converts a Unix time to byte slice
func EpochToByteSlice(timestamp int64) []byte {
	return []byte(strconv.FormatInt(timestamp, 10))
}

//UintToByteSlice converts int64 to byte slices
func UintToByteSlice(num uint64) []byte {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, num)
	return buffer
}
