package main

import "strconv"

func hammingWeight(num uint32) int {
	var sum int
	numStr := strconv.FormatUint(uint64(num), 2)
	for i := 0; i < len(numStr); i++ {
		if numStr[i]-'0' == 1 {
			sum++
		}
	}
	return sum
}
