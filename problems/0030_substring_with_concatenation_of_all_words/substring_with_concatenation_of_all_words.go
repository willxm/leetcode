package problem

func findSubstring(s string, words []string) []int {
	var res []int
	var allIndex [][]int
	wl := len(words[0])

	for _, w := range words {
		indexs := containsIndex(s, w)
		allIndex = append(allIndex, indexs)
	}

	for i := 0; i < len(allIndex); i++ {
		for j := 0; j < len(allIndex[i]); j++ {
			// if abs(allIndex[i])
		}
	}

	return res
}

func containsIndex(s, ss string) []int {
	l := len(ss)
	rs := []rune(s)
	rss := []rune(ss)
	var res []int
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(rss); j++ {
			if rs[i] == rss[j] {
				j++
			} else {
				i++
				j = 0
			}
		}
		res = append(res, i-l)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
