package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	var sum uint32
	numStr := fmt.Sprintf("%0*v", 32, strconv.FormatUint(uint64(num), 2))
	for i := len(numStr) - 1; i >= 0; i-- {
		sum = sum*2 + uint32(numStr[i]-'0')
	}
	return sum
}
