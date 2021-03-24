package problems

func mySqrt(x int) int {
	res := x

	//Newton's method
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
