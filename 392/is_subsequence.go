package main

func isSubsequence(s string, t string) bool {
	p := 0
	for i := 0; i < len(t) && p < len(s); i++ {
		if s[p] == t[i] {
			p++
		}
	}

	return p == len(s)
}

func isSubsequence2(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen > tLen {
		return false
	}

	if sLen == 0 {
		return true
	}

	a := [100][10001]bool{}
	for i := 0; i <= tLen; i++ {
		a[0][i] = true
	}
	for i := 1; i <= sLen; i++ {
		a[i][0] = false
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if s[i-1] == t[j-1] {
				a[i][j] = a[i-1][j-1]
			} else {
				a[i][j] = a[i][j-1]
			}
		}
	}

	return a[sLen][tLen]
}
