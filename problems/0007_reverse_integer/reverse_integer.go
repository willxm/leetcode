package problems

func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}
	if x == 0 {
		return 0
	}
	ret := 0

	for x != 0 {
		ret = ret*10 + x%10
		if ret < -2147483648 || ret > 2147483647 {
			return 0
		}
		x = x / 10
	}
	return ret
}
