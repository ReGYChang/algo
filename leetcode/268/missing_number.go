package main

import "sort"

// map
func missingNumber(nums []int) int {
	l := len(nums)
	m := make(map[int]int, l)
	for _, val := range nums {
		m[val] = val
	}
	for i := 0; i < l+1; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0
}

// sorting slice
func missingNumber2(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		if n := nums[i]; i < n {
			return i
		}
	}
	return l
}
