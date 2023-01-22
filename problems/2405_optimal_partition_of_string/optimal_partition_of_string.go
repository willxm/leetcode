package main

func partitionString(s string) int {
	rs := []rune(s)
	tempM := 0

	ans := 1
	for i := 0; i < len(rs); i++ {
		l := rs[i] - 97
		if tempM&(1<<l) == 0 {
			tempM |= 1 << l
		} else {
			ans++
			tempM = 0
			tempM |= 1 << l
		}
	}
	return ans
}
