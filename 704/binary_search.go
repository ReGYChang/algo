package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		n := nums[mid]

		if n > target {
			right = mid - 1
		} else if n < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
