package main

func plusOne(digits []int) []int {
	s := make([]int, len(digits)+1)
	pre := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if res := digits[i] + pre; res >= 10 {
			pre = 1
			s[i+1] = res - 10
		} else {
			pre = 0
			s[i+1] = res
		}
	}

	if pre == 0 {
		return s[1:]
	}
	s[0] = 1
	return s
}
