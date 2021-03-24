package problems

func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	minP := prices[0]

	profit := prices[1] - prices[0]

	for i := 2; i < l; i++ {
		minP = min(minP, prices[i-1])
		profit = max(profit, prices[i]-minP)
	}

	if profit < 0 {
		return 0
	}
	return profit
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
