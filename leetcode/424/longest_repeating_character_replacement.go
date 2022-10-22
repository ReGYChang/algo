package main

func characterReplacement(s string, k int) int {
	left, right, max, total, freq := 0, 0, 0, 0, [26]int{}
	for ; right < len(s); right++ {
		freq[s[right]-'A']++
		if freq[s[right]-'A'] > max {
			max = freq[s[right]-'A']
		}

		if total-max >= k {
			left++
			freq[s[left-1]-'A']--
		} else {
			total++
		}
	}

	return right - left
}
