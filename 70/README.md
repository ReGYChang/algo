- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Constraints](#constraints)
- [Approach 1: Dynamic Programming](#approach-1-dynamic-programming)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Example 1

- Input: n = 2
- Output: 2
- Explanation: There are two ways to climb to the top.
  - 1 step + 1 step
  - 2 steps


## Example 2

- Input: n = 3
- Output: 3
- Explanation: There are three ways to climb to the top.
  - 1 step + 1 step + 1 step
  - 1 step + 2 steps
  - 2 steps + 1 step

## Constraints

- 1 <= n <= 45

# Approach 1: Dynamic Programming

## Algorithm

爬梯子過程中每次只能向上爬一步或是向上爬兩步, 即上到當前高度有兩種方法, 一種為當前高度下一階, 一種為當前高度下兩階

假設 $C(n)$ 表示爬到高度 `n` 的不同方法總數, 則 $C(n)$ = $C(n-1)$ + $C(n-2)$

## Implementation

```go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1] + dp[i-2]
	}

	return dp[n]
}
```

## Complexity Analysis

- Time complexity: O($n$)
  - Where `n` is the intergers given.
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.

- Space complexity: O($n$)
  - Where `n` is the intergers given.
  - Memory Usage: 1.9 MB, less than 91.21% of Go online submissions for Climbing Stairs.
