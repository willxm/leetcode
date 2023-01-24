package main

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = n
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
