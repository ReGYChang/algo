package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
	l := len(s)
	for idx, val := range s {
		if byte(val) != s[l - idx - 1] {
			return false
		}
	}
	return true
}