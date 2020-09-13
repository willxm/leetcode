package problem

func minCostConnectPoints(points [][]int) int {
	l := len(points)

	res := make([][]int, l)
	for i := 0; i < l; i++ {
		res[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			d := abs((points[j][0] - points[i][0])) + abs((points[j][1] - points[i][1]))
			res[i][j] = d
		}
	}

	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
