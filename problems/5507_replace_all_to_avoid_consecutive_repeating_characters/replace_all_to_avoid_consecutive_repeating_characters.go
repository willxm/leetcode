package problems

func modifyString(s string) string {
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] == '?' {
			ss[i] = 'a'
			if i-1 >= 0 && ss[i] == ss[i-1] {
				ss[i] = ss[i] + 1
			}
			if i+1 < len(ss) && ss[i] == ss[i+1] {
				ss[i] = ss[i] + 1
			}
			if i-1 >= 0 && ss[i] == ss[i-1] {
				ss[i] = ss[i] + 1
			}
		}
	}
	return string(ss)
}
