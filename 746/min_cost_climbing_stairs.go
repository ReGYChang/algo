package main

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
