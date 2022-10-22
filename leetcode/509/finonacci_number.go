package main

func fibRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}
