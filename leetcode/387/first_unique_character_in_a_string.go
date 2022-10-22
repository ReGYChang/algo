package main

func firstUniqChar(s string) int {
	var freq [26]int
	l := len(s)
	for i := 0; i < l; i++ {
		freq[s[i]-'a']++
	}
	for i := 0; i < l; i++ {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
