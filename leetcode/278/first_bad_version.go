package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var Bad int

func firstBadVersion(n int) int {
	base := 1
	for base <= n {
		mid := (base + n) / 2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			n = mid - 1
		} else {
			base = mid + 1
		}
	}
	return 1
}

func isBadVersion(version int) bool {
	return version == Bad
}
