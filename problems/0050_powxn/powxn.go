package problems

func myPow(x float64, n int) float64 {
	f := 0
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		f = 1
	}
	res := myPow(x, n>>1)
	res *= res
	if n&0x1 == 1 {
		res *= x
	}
	if f == 1 {
		return 1.0 / res
	}
	return res
}
