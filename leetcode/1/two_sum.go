package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		if n, ok := m[target-val]; ok {
			return []int{n, idx}
		} else {
			m[val] = idx
		}
	}
	return nil
}
