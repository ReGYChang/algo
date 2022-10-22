- [Problem](#problem)
	- [Example 1:](#example-1)
	- [Example 2:](#example-2)
	- [Constraints:](#constraints)
- [Approach 1: Array](#approach-1-array)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)

# Problem

You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

- The number of "bulls", which are **digits in the guess that are in the correct position**.
- The number of "cows", which are **digits in the guess that are in your secret number but are located in the wrong position**. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.

Given the secret number `secret` and your friend's guess `guess`, return the hint for your friend's guess.

The hint should be formatted as `"xAyB"`, where `x` is the number of bulls and `y` is the number of cows. Note that both secret and guess may contain duplicate digits.

## Example 1:

```
- Input: secret = "1807", guess = "7810"
- Output: "1A3B"
- Explanation: Bulls are connected with a '|' and cows are underlined:

"1807"
  |
"7810"
```

## Example 2:

```
- Input: secret = "1123", guess = "0111"
- Output: "1A1B"
- Explanation: Bulls are connected with a '|' and cows are underlined:

"1123"        "1123"
  |      or     |
"0111"        "0111"

Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.
```

## Constraints:

- 1 <= secret.length, guess.length <= 1000
- secret.length == guess.length
- secret and guess consist of digits only.

# Approach 1: Array

## Algorithm

- 若 `secret` 與 `guess` 在相同位置上的數字相同則算一個 `bulls`
- 若 `secret` 與 `guess` 在相同位置上的數字不同, 但在不同位置上有相同數字, 則算一個 `cows`

此題使用 `array` 的方式紀錄每個 char 出現頻率, 排除相同位置相同 char, 再判斷是否有相同數字即可

## Implementation

```go
func getHint(secret string, guess string) string {
	a, b, l := 0, 0, len(secret)
	sFreq := [10]int{}
	gFreq := [10]int{}

	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			sFreq[secret[i]-'0']++
			gFreq[guess[i]-'0']++
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
```

## Complexity Analysis

- Time complexity: O(n)

  - Where `n` is the length of given string.
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for Bulls and Cows.


- Space complexity: O(1)

  - Only constant space is used.
  - Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Bulls and Cows.
