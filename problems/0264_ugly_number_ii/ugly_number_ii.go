package problem

func nthUglyNumber(n int) int {
	var us = []int{2, 3, 4, 5}
	if n <= 0 {
		return 0
	}
	if n < 5 {
		return us[n]
	}

	number := 4

	for i := 5; i <= n; i++ {
		min(us[number-1])
	}

}

func min(x, y, z int) int {
	min := 0
	if x < y {
		min = x
	} else {
		min = y
	}

	if min < z {
		min = min
	} else {
		min = z
	}
	return min
}
