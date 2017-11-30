package main

import (
	// "bytes"
	// "encoding/binary"
	// "log"
	"strconv"
)

// IntToHex并不是已有包中自带函数，需要自己实现
// IntToHex convert int to hexadecimal representation
func IntToHex(i int64) []byte {
	// return fmt.Sprintf("0x%x", i)
	return []byte(strconv.FormatInt(i, 10))
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
