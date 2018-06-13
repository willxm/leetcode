package problems

import (
	"math"
)

func longestPalindrome(s string) string {
	lenStr := len(s)

	begin, maxLen := 0, 1

	if lenStr < 2 {
		return s
	}

	for i := 0; i < lenStr; i++ {
		for j := lenStr; j > i; j-- {
			if maxLen > j-i {
				break
			}
			if isPalindromic(s[i:j]) {
				if j-i > maxLen {
					begin = i
					maxLen = j - i
				}
			}
		}
	}
	return s[begin : begin+maxLen]
}

func isPalindromic(s string) bool {
	lenStr := len(s)
	if lenStr == 1 || lenStr == 0 {
		return true
	}
	mid := int(math.Ceil(float64(lenStr) / 2.0))
	for i := 0; i < lenStr; i++ {
		if s[i:i+1] == s[lenStr-i-1:lenStr-i] {
			if i+1 == mid {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
