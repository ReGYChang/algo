package main

func mySqrt(x int) int {
	var l, r = 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l
}
