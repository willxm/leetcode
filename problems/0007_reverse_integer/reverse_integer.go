package problem

import (
	"strconv"
)

func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}
	if x < 0 {
		x = -x
		x = deleteZero(x)
		s := reverseString(strconv.Itoa(x))
		res, _ := strconv.Atoi(s)
		if -res < -2147483648 || -res > 2147483647 {
			return 0
		}
		return -res
	} else if x > 0 {
		x = deleteZero(x)
		s := reverseString(strconv.Itoa(x))
		res, _ := strconv.Atoi(s)
		if res < -2147483648 || res > 2147483647 {
			return 0
		}
		return res
	} else {
		return 0
	}
	return 0
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func deleteZero(x int) int {
	for x%10 == 0 {
		x = x / 10
	}
	return x
}
