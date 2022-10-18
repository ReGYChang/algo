package main

func isAnagram(s string, t string) bool {
	var freq [128]int
	var sLen, tLen = len(s), len(t)
	for i := 0; i < tLen; i++ {
		freq[t[i]]++
	}
	for i := 0; i < sLen; i++ {
		if freq[s[i]] == 0 {
			return false
		} else {
			freq[s[i]]--
		}
	}
	for i := 0; i < 128; i++ {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}
