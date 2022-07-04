package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	suffix := make([]int, l)
	res := make([]int, l)
	prefix[0] = 1
	suffix[l-1] = 1

	for i := 1; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := l - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < l; i++ {
		res[i] = prefix[i] * suffix[i]
	}

	return res
}
