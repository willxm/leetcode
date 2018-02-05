package problems

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	lenStr := len(s)

	// maxSub := ""
	if lenStr == 0 {
		return 0
	}
	maxLen := 1

	for i := 0; i < lenStr; i++ {
		sub := s[i : i+1]
		l := 1
		for j := i + 1; j < lenStr; j++ {
			if strings.Contains(sub, s[j:j+1]) {
				l = 1
				break
			} else {
				sub += s[j : j+1]
				l++
			}

			if l > maxLen {
				maxLen = l
				// maxSub = sub
			}
		}
	}
	return maxLen
}
