package problems

import "fmt"

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = triangle[l-1][i]
	}
	fmt.Println(dp)
	for i := l - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
			fmt.Println(dp, i, j)
		}
	}
	return dp[0]
}

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}
