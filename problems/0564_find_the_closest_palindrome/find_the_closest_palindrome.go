package problem

func nearestPalindromic(n string) string {
	runes := []rune(n)
	l := len(runes)
	mid := int((l - 1) / 2)

	for i := mid + 1; i < l; i++ {
		runes[i] = runes[l-i-1]
	}
	oT := runes
	res := string(oT)

	if res == n {
		if len(n)%2 == 0 {

		} else {

		}
	}
	return res
}
