package main

func removeDuplicates(nums []int) int {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[pos] != nums[i] {
			pos++
			nums[pos] = nums[i]
		}
	}

	return pos + 1
}

func removeDuplicatesBubble(nums []int) int {
	var freq [201]bool
	var numsLen = len(nums)
	for i := 0; i < numsLen; {
		if freq[nums[i]+100] {
			numsLen--
			sink(i, nums)
		} else {
			freq[nums[i]+100] = true
			i++
		}
	}

	return numsLen
}

func sink(i int, nums []int) {
	l := len(nums)
	for j := i + 1; j < l; j++ {
		swap(i, j, nums)
		i++
	}
}

func swap(i, j int, s []int) {
	s[i], s[j] = s[j], s[i]
}
