- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Example 3](#example-3)
  - [Constraints](#constraints)
- [Approach 1: Recursive](#approach-1-recursive)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

```go
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
```

Given `n`, calculate `F(n)`.


## Example 1

- Input: n = 2
- Output: 1
- Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

## Example 2

- Input: n = 3
- Output: 2
- Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

## Example 3

- Input: n = 4
- Output: 3
- Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

## Constraints

- 0 <= n <= 30

# Approach 1: Recursive

## Algorithm

對於任意大於一的正整數而言, `F(n)` = `F(n-1)` + `F(n-2)`, 可直接透過 recursive 的方式求解

## Implementation

```go
func fibRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}
```


## Complexity Analysis

- Time complexity: $O(2^n)$
  - Recursion tree which will have depth `n` and intuitively figure out that this function is asymptotically $O(2^n)$
  - Runtime: 13 ms, faster than 21.35% of Go online submissions for Fibonacci Number.

- Space complexity: O(n)
  - Recursion tree which will have depth `n`, for function call stack.
  - Memory Usage: 6.2 MB, less than 34.38% of Go online submissions for N-ary Tree Preorder Traversal.