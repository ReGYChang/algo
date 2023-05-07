# Problem

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

## Example 1

- Input: digits = "23"
- Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]


## Example 2

- Input: digits = ""
- Output: []

## Example 3

- Input: digits = "2"
- Output: ["a","b","c"]

## Constraints

- 0 <= digits.length <= 4
- digits[i] is a digit in the range ['2', '9'].

# Approach 1: Backtracking

## Algorithm



## Implementation

```go
var m = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}

	letters := m[digits[0]]
	lastCombinations := letterCombinations(digits[1:])
	for i := 0; i < len(letters); i++ {
		if len(lastCombinations) == 0 {
			res = append(res, letters[i])
		} else {
			for j := 0; j < len(lastCombinations); j++ {
				res = append(res, letters[i]+lastCombinations[j])
			}
		}
	}

	return res
}
```

## Complexity Analysis

### Time complexity: $O(4^n)$

1. The function `letterCombinations` is called recursively, and each recursive call reduces the length of the input string `digits` by 1. This means that the recursion depth is equal to the length of the input string `digits`, which we can denote as n.

2. For each level of the recursion, the code iterates through the possible letters of the current digit (using the `for i` loop) and combines them with the combinations returned by the recursive call (using the `for j` loop). In the worst case, a digit maps to 4 characters (e.g., '7' or '9'), so the maximum number of iterations for both loops is 4.

3. The cost of the string concatenation operation (`letters[i] + lastCombinations[j]`) can be assumed to be constant time, as the length of the concatenated strings is equal to the recursion depth, which is at most n.

Considering the points above, the time complexity of the algorithm can be represented as $T(n) = 4 * T(n-1)$ with the base case $T(0) = 1$. This recurrence relation can be solved to get $T(n) = 4^n$, which means that the time complexity of the function is $O(4^n)$.

> Runtime: 0 ms, faster than 100.00% of Go online submissions for  Letter Combinations of a Phone Number.

### Space complexity: $O(4^n * n)$

1. The space complexity is affected by the depth of the recursion, the number of function calls, and the size of the result array res.
   
2. The depth of the recursion is n, as mentioned earlier.
   
3. In each level of the recursion, the size of the result array increases. In the worst case, each digit maps to 4 characters, and since we're forming combinations of all characters, the size of the result array can grow up to 4^n.
   
4. Additionally, each string in the result array has a length at most n, which also contributes to the space complexity.

Considering the points above, the space complexity of the function can be represented as $S(n) = 4^n * n$. This means that the space complexity of the function is $O(4^n * n)$.

> Memory Usage: 2 MB, less than 78.41% of Go online submissions for  Letter Combinations of a Phone Number.