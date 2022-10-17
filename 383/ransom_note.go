package main

func canConstruct(ransomNote string, magazine string) bool {
	var freq [128]int
	var rLen, mLen = len(ransomNote), len(magazine)
	for i := 0; i < mLen; i++ {
		freq[magazine[i]]++
	}
	for i := 0; i < rLen; i++ {
		if freq[ransomNote[i]] == 0 {
			return false
		} else {
			freq[ransomNote[i]]--
		}
	}
	return true
}
