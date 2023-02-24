package main

func countBits(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		n := i
		c := 0
		for n > 0 {
			if n&1 == 1 {
				c++
			}
			n >>= 1
		}
		res = append(res, c)
	}
	return res
}
