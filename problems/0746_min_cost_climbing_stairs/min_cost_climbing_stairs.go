package main

// func minCostClimbingStairs(cost []int) int {
// 	return dp(len(cost), cost)
// }

// func dp(n int, cost []int) int {
// 	if n <= 1 {
// 		return 0
// 	}
// 	return min(dp(n-1, cost)+cost[n-1], dp(n-2, cost)+cost[n-2])
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
