package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(dp(0, n-2, nums), dp(1, n-1, nums))
}

func dp(i, n int, nums []int) int {
	if n < 0 {
		return 0
	}
	if n < i {
		return 0
	}
	return max(dp(i, n-2, nums)+nums[n], dp(i, n-1, nums))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
