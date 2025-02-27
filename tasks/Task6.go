package tasks

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Task6() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}
