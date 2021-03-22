package problems

import "strings"

func myAtoi(s string) int {
	ret := 0
	symbol := 0

	s = strings.Trim(s, " ")
	rs := []rune(s)
	if len(rs) == 0 {
		return 0
	}

	for i := 0; i < len(rs); i++ {
		// if rs[i] == ' ' {
		// 	continue
		// }

		if ret == 0 {
			if rs[i] == '+' || rs[i] == '-' {
				if i+1 < len(rs) && (rs[i+1] == '+' || rs[i+1] == '-' || rs[i+1] == ' ') {
					ret = 0
					break
				}
			}

			if rs[i] == '0' {
				if i+1 < len(rs) && (rs[i+1] == '+' || rs[i+1] == '-' || rs[i+1] == ' ') {
					ret = 0
					break
				}
			}
		}

		if rs[i] >= '0' && rs[i] <= '9' {
			if i != 0 && rs[i-1] == '-' {
				symbol = 1
			}
			ret = ret*10 + int(rs[i]-'0')
			if ret > 2147483647 && symbol == 1 {
				ret = 2147483648
				break
			}

			if ret > 2147483647 && symbol == 0 {
				ret = 2147483647
				break
			}

		} else {
			if ret > 0 {
				break
			}
			if rs[i] != ' ' && rs[i] != '-' && rs[i] != '+' {
				ret = 0
				break
			}
		}
	}
	if symbol == 1 {
		return 0 - ret

	}
	return ret
}
