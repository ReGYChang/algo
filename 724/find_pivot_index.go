package main

func pivotIndex(nums []int) int {
	for idx, _ := range nums {
		var leftSum, rightSum int
		for _, val := range nums[:idx] {
			leftSum += val
		}
		for _, val := range nums[idx+1:] {
			rightSum += val
		}
		if leftSum == rightSum {
			return idx
		}
	}
	return -1
}

func pivotIndex2(nums []int) int {
	l := len(nums)
	s := make([]int, l) // slice s record left sum
	sum := 0
	for i := 1; i < l; i++ {
		s[i] += s[i-1] + nums[i-1]
		sum += nums[i]
	}
	sum += nums[0]
	for i := 0; i < l; i++ {
		if s[i] == sum - s[i] - nums[i] {
			return i
		}
	}
	return -1
}

func pivotIndex3(nums []int) int {
	l := len(nums)
	sum, leftSum := 0, 0
	for _, val := range nums {
		sum += val
	}
	for i := 0; i < l; i++ {
		if leftSum == sum - leftSum - nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}