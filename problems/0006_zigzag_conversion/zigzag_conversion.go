package problems

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	runes := []rune(s)
	rows := make([][]rune, numRows)
	i := 0
	flag := 0
	for index := range runes {
		rows[i] = append(rows[i], runes[index])
		if flag == 0 {
			i++
			if i == numRows-1 {
				flag = 1
			}
		} else {
			i--
			if i == 0 {
				flag = 0
			}
		}
	}
	retRunes := []rune{}
	for index := range rows {
		retRunes = append(retRunes, rows[index]...)
	}
	return string(retRunes)
}
