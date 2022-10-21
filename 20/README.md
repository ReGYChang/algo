- [Problem](#problem)
	- [Example 1](#example-1)
	- [Example 2](#example-2)
	- [Example 3](#example-3)
	- [Constraints](#constraints)
- [Approach 1: Iteratively](#approach-1-iteratively)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)

# Problem

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

## Example 1

- Input: s = "()"
- Output: true


## Example 2

- Input: s = "()[]{}"
- Output: true

## Example 3

- Input: s = "(]"
- Output: false

## Constraints

- 1 <= s.length <= 104
- s consists of parentheses only '()[]{}'.

# Approach 1: Iteratively

## Algorithm

透過 stack 來確認字串是否有效, 若遇到 `(`, `{` `[` 時將字元 push 到 stack, 否則 pop element from stack 並檢查是否匹配

需注意 cornar cases, 如 `(`, `)` 等

## Implementation

```go
func isValid(s string) bool {
	var stack []uint8
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				switch s[i] {
				case ')':
					if stack[len(stack)-1] != '(' {
						return false
					}
					stack = stack[:len(stack)-1]
				case ']':
					if stack[len(stack)-1] != '[' {
						return false
					}
					stack = stack[:len(stack)-1]
				case '}':
					if stack[len(stack)-1] != '{' {
						return false
					}
          stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
```

## Complexity Analysis

- Time complexity: O(n)
    - Where `n` is the number of character in the given string.
    - Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.

- Space complexity: O(n)
    - Where `n` is the number of character in the given string.
    - Memory Usage: 1.9 MB, less than 99.53% of Go online submissions for Valid Parentheses.
