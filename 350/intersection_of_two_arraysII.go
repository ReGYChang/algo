package main

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]]++
	}

	res := make([]int, 0)
	for k, v := range m1 {
		if v2, exist := m2[k]; exist {
			for i := 0; i < min(v, v2); i++ {
				res = append(res, k)
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
