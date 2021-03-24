package problems

func maxDepth(s string) int {
	rs := []rune(s)

	maxLen := 0
	tmpLen := 0

	for _, v := range rs {
		if v == '(' {
			tmpLen++
		}

		if v == ')' {
			tmpLen--
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}
	return maxLen
}
