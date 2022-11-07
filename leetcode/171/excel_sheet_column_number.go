package main

func titleToNumber(columnTitle string) int {
	var sum int
	for i := 0; i < len(columnTitle); i++ {
		sum = sum*26 + int(columnTitle[i]-'A'+1)
	}
	return sum
}
