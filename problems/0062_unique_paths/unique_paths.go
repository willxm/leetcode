package problem

func uniquePaths(m int, n int) int {
	sum := m + n
	m = max(m, n)
	n = sum - m
	return c(m-1+n-1, n-1)
}

func c(m, n int) int {
	x := 1
	y := 1
	for i := 0; i < n; i++ {
		x *= m - i
		y *= i + 1
	}
	return x / y
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
