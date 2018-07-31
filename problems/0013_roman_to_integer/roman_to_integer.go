package problem

type roman struct {
	symbol string
	value  int
}

var romansSingle = []roman{
	{"M", 1000},
	{"D", 500},
	{"C", 100},
	{"L", 50},
	{"X", 10},
	{"V", 5},
	{"I", 1},
}

var romansDuo = []roman{
	{"CM", 900},
	{"CD", 400},
	{"XC", 90},
	{"XL", 40},
	{"IX", 9},
	{"IV", 4},
}

func romanToInt(s string) int {
	var res int
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for k, v := range romansSingle {
			if string(runes[i]) == v.symbol {
				if i < len(runes)-1 && k > 1 {
					if string(runes[i+1]) == romansSingle[k-1].symbol {
						res += romansSingle[k-1].value - romansSingle[k].value
						i++
						break
					} else if string(runes[i+1]) == romansSingle[k-2].symbol {
						res += romansSingle[k-2].value - romansSingle[k].value
						i++
						break
					} else {
						res += v.value
						break
					}
				} else {
					res += v.value
					break
				}

			}
		}
	}
	return res
}
