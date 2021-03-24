package problems

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	minStr := []rune(getShortestStr(strs))

	var res string
	tempPrefix := ""

	for _, c := range minStr {
		tempPrefix += string(c)
		k := 0
		for _, v := range strs {
			if strings.HasPrefix(v, tempPrefix) {
				k++
			}
		}
		if k == len(strs) {
			res = tempPrefix
		} else {
			return res
		}
	}
	return res

}

func getShortestStr(strs []string) string {
	minStr := 99999
	minIndex := -1
	for k, v := range strs {
		if len(v) < minStr {
			minStr = len(v)
			minIndex = k
		}
	}
	return strs[minIndex]
}
