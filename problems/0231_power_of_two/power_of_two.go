package problems

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	hasOne := false
	for n > 0 {
		if (n & 1) != 0 {
			if hasOne {
				return false
			} else {
				hasOne = true
			}
		}
		n >>= 1
	}
	return hasOne
}
