package main

func numberOfSteps(num int) int {
	ans := 0
	for {
		if num == 0 {
			return ans
		} else if num&1 == 0 {
			// num /= 2
			num = num >> 1
		} else {
			num--
		}
		ans++
	}
}
