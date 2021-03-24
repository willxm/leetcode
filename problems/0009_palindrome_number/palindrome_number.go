package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		tmp := x
		y := 0

		for x != 0 {
			y = y*10 + x%10
			x = x / 10
		}
		if y == tmp {
			return true
		} else {
			return false
		}
	}
}
