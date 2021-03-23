package problems

func isMatch(s string, p string) bool {
	rs := []rune(s)
	rp := []rune(p)
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(rp); j++ {
			if rs[i] == rp[j] {
				i++
			} else {

			}
		}
	}
	return true
}
