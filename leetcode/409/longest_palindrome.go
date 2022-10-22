package main

func longestPalindrome(s string) int {
	m, l, res, oddExist := make(map[uint8]int), len(s), 0, false
	for i := 0; i < l; i++ {
		m[s[i]]++
	}

	for _, v := range m {
		if v&1 == 0 {
			res += v
		} else {
			res += v - 1
			if !oddExist {
				oddExist = true
			}
		}
	}

	if oddExist {
		return res + 1
	} else {
		return res
	}
}
