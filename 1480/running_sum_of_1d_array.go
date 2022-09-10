package main

func runningSum(nums []int) []int {
	sum := 0
	result := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = sum + nums[i]
		sum += nums[i]
	}

	return result
}
