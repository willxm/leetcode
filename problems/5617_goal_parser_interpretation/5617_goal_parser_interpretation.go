package problem

func interpret(command string) string {
	var res string
	rs := []rune(command)
	for i := 0; i < len(rs); i++ {
		if rs[i] == '(' {
			if rs[i+1] == ')' {
				res += "o"
				i = i + 1
			} else {
				res += "al"
				i = i + 3
			}
		} else {
			res += string(rs[i])
		}
	}
	return res
}
