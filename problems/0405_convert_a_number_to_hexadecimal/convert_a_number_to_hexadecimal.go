package problems

import "strings"

var (
	hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
)

func toHex(num int) string {
	if num < 0 {
		num = 4294967296 + num
	}
	var res []string
	for num/16 > 0 {
		res = append(res, hex[num%16])
		num /= 16
	}
	res = append(res, hex[num%16])
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return strings.Join(res, "")
}
