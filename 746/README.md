- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Constraints](#constraints)
- [Approach 1: Dynamic Programming](#approach-1-dynamic-programming)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

You are given an integer array cost where `cost[i]` is the cost of ith step on a staircase. Once you pay the cost, you can either climb `one` or `two` steps.

You can either start from the step with index `0`, or the step with index `1`.

Return the `minimum` cost to reach the top of the floor.


## Example 1

- Input: cost = [10,15,20]
- Output: 15
- Explanation: 
  - You will start at index 1.
  - Pay 15 and climb two steps to reach the top.
  - The total cost is 15.


## Example 2

- Input: cost = [1,100,1,1,1,100,1,1,100,1]
- Output: 6
- Explanation: 
  - You will start at index 0.
  - Pay 1 and climb two steps to reach index 2.
  - Pay 1 and climb two steps to reach index 4.
  - Pay 1 and climb two steps to reach index 6.
  - Pay 1 and climb one step to reach index 7.
  - Pay 1 and climb two steps to reach index 9.
  - Pay 1 and climb one step to reach the top.
  - The total cost is 6.

## Constraints

- 2 <= `cost.length` <= 1000
- 0 <= `cost[i]` <= 999

# Approach 1: Dynamic Programming

## Algorithm

可創建一個 slice `dp`, `dp[i]` 表示爬到第 `i` 階所需要花費的最少步數, 為前一階或是前兩階的最小步數加上當前階層所需步數, 即:

> dp[i] = min(dp[i-2], dp[i-1]) + cost[i]

## Implementation

```go
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l, l)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < l; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[l-1], dp[l-2])
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
```

## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the length of given slice.
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.

- Space complexity: $O(m*n)$
  - Where `n` is the length of given slice.
  - Memory Usage: 3.1 MB, less than 43.33% of Go online submissions for Min Cost Climbing Stairs.
