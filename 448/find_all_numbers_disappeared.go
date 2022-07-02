package main

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, val := range nums {
		m[val] = val
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

func findDisappearedNumbers2(nums []int) []int {
	l := len(nums)
	s := make([]int, l)
	res := make([]int, 0)
	for _, val := range nums {
		s[val-1] = val
	}
	for idx, val := range s {
		if val == 0 {
			res = append(res, idx+1)
		}
	}
	return res
}
