package problem

func minCost(s string, cost []int) int {
	ss := []rune(s)
	c := 0
	for i := 0; i < len(ss); i++ {
		if i+1 < len(ss) && ss[i] == ss[i+1] {
			m, im := min(cost[i], cost[i+1])
			c += m
			if im == 0 {
				cost[i] = cost[i+1]
			} else {
				cost[i+1] = cost[i]
			}
		}
	}
	return c
}

func min(a, b int) (int, int) {
	if a < b {
		return a, 0
	}
	return b, 1
}
