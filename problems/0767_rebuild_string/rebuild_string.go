package string

func reorganizeString(S string) string {
	rs := []rune(S)
	// start := 0
	end := len(rs) - 1

	i := 1
	j := end
	for ; i <= end; i++ {
		for ; j > i; j-- {

			if rs[i-1] == rs[i] {
				if rs[i] == rs[j] {
					break
				}
				rs[i], rs[j] = rs[j], rs[i]
				// i++
				// j = end
				break

			} else {
				break
			}
		}
		if i == j {
			return ""
		}

	}
	return string(rs)
}
