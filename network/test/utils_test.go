package test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	bytes := make([]byte, 12, 12)
	fmt.Printf("bytes = %v\n", bytes)
	bytes[0] = 0
	bytes[1] = 1
	bytes[2] = 2
	bytes[3] = 3
	binary.LittleEndian.PutUint32(bytes[4:], 111)
	binary.LittleEndian.PutUint32(bytes[8:], 2222222)
	fmt.Printf("bytes = %v\n", bytes)
	bytes = bytes[12:]
	fmt.Printf("bytes = %v\n", bytes)

}
