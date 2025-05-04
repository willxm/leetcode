package problems

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	nagative := (dividend < 0) != (divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var result int
	for dividend >= divisor {
		dividend = dividend - divisor
		result++
	}
	if nagative {
		result = -result
	}
	return result
}
