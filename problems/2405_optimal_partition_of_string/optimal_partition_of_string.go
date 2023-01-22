package main

func partitionString(s string) int {
	rs := []rune(s)
	tempM := make([]byte, 26)

	ans := 1
	for i := 0; i < len(rs); i++ {
		if tempM[rs[i]-97] == 0 {
			tempM[rs[i]-97] = 1
		} else {
			ans++
			tempM = make([]byte, 26)
			tempM[rs[i]-97] = 1
		}
	}
	return ans
}
