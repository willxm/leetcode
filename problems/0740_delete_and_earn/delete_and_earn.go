package main

func deleteAndEarn(nums []int) int {
	houses := make([]int, 10001)
	for _, v := range nums {
		houses[v] += v
	}
	n := len(houses)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = houses[0]
	for i := 2; i < n+1; i++ {
		dp[i] = max(dp[i-2]+houses[i-1], dp[i-1])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
