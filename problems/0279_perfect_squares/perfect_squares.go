package problem

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = 999999999
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			dp[i+j*j] = min(dp[i+j*j], dp[i]+1)

		}
	}
	return dp[n]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
