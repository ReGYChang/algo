package main

func findAnagrams(s string, p string) []int {
	var sl, pl = len(s), len(p)
	var freq, pFreq [256]int
	res := make([]int, 0)
	if sl < pl {
		return res
	}

	for i := 0; i < pl; i++ {
		pFreq[p[i]] += 1
		freq[s[i]] += 1
	}
	if freq == pFreq {
		res = append(res, 0)
	}
	for i := 1; i < sl-pl+1; i++ {
		freq[s[i-1]]--
		freq[s[i+pl-1]]++

		if freq == pFreq {
			res = append(res, i)
		}
	}

	return res
}
