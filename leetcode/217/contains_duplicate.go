package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		} else {
			m[val] = val
		}
	}
	return false
}
