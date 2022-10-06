package main

import "fmt"

func getHint(secret string, guess string) string {
	a, b, l := 0, 0, len(secret)
	sFreq := [10]int{}
	for i := 0; i < l; i++ {
		sFreq[secret[i]-'0']++
	}
	gFreq := [10]int{}
	for i := 0; i < l; i++ {
		gFreq[guess[i]-'0']++
	}

	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			a++
			sFreq[secret[i]-'0']--
			gFreq[guess[i]-'0']--
		}
	}
	for i := 0; i < 10; i++ {
		if sFreq[i] > 0 && gFreq[i] > 0 {
			b += min(sFreq[i], gFreq[i])
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
