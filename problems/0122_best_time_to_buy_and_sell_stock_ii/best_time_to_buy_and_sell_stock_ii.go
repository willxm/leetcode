package problem

func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	sum := 0

	for i := 1; i < l; i++ {

		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
