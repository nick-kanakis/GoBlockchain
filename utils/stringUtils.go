package utils

import "bytes"

//ConcatStrings concatenate strings in O(n)
func ConcatStrings(strings ...string) string {
	var buffer bytes.Buffer

	for _, s := range strings {
		buffer.WriteString(s)
	}
	return buffer.String()
}
