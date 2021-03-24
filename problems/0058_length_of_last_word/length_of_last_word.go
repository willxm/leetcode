package problems

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	l := len(s)
	if l == 0 {
		return 0
	}
	ss := strings.Split(s, " ")
	return len(ss[len(ss)-1])
}
