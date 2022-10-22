package main

func singleNumber(nums []int) int {
	n := nums[0]
	l := len(nums)
	for i := 1; i < l; i++ {
		n = n ^ nums[i]
	}
	return n
}
