package main

// func rob(nums []int) int {
// 	return dp(len(nums)-1, nums)
// }

// func dp(n int, nums []int) int {
// 	if n < 0 {
// 		return 0
// 	}
// 	return max(dp(n-2, nums)+nums[n], dp(n-1, nums))
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < n+1; i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[n]
}
