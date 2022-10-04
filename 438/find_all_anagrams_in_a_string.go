package main

func findAnagrams(s string, p string) []int {
	var freq, pFreq [26]int
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}

	for i := 0; i < len(p); i++ {
		pFreq[p[i]-'a'] += 1
		freq[s[i]-'a'] += 1
	}
	if freq == pFreq {
		res = append(res, 0)
	}
	for i := 1; i < len(s)-len(p)+1; i++ {
		freq[s[i-1]-'a']--
		freq[s[i+len(p)-1]-'a']++

		if freq == pFreq {
			res = append(res, i)
		}
	}

	return res
}
