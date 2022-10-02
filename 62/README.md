- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Constraints](#constraints)
- [Approach 1: Dynamic Programming](#approach-1-dynamic-programming)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

There is a robot on an m x n grid. The robot is initially located at the `top-left corner` (i.e., `grid[0][0]`). The robot tries to move to the `bottom-right corner` (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.


## Example 1

- Input: m = 3, n = 7
- Output: 28


## Example 2

- Input: m = 3, n = 2
- Output: 3
- Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
  - Right -> Down -> Down
  - Down -> Down -> Right
  - Down -> Right -> Down

## Constraints

- 1 <= m, n <= 100

# Approach 1: Dynamic Programming

## Algorithm

可創建一個 $m*n$ 的 slice `dp`, `dp[m][n]` 表示機器人走到 `grid[m][n]` 所需要走的步數

機器人只能向又或向下走, 因此走到 `grid[m][n]` 所需的步數為 `dp[m-1][n]` + `dp[m][n-1]`

## Implementation

```go
func uniquePaths(m int, n int) int {
	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
	}

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
```

## Complexity Analysis

- Time complexity: $O(m*n)$
  - Where `m` and `n` is the intergers given.
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.

- Space complexity: $O(m*n)$
  - Where `m` and `n` is the intergers given.
  - Memory Usage: 1.9 MB, less than 75.66% of Go online submissions for Unique Paths.
