package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var maxKey, maxVal int
	for k, v := range m {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}

	return maxKey
}
