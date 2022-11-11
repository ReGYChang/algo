package main

import (
	"strconv"
)

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		s := strconv.Itoa(n)
		n = 0
		for i := 0; i < len(s); i++ {
			n += int(s[i]-'0') * int(s[i]-'0')
		}

		if _, exist := m[n]; exist {
			return false
		}
		m[n] = true
	}
	return true
}
