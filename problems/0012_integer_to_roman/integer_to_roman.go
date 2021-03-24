package problems

type roman struct {
	symbol string
	value  int
}

var romans = []roman{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func intToRoman(num int) string {
	var res string
	for _, v := range romans {
		flag := num / v.value
		if flag > 0 {
			for i := 0; i < flag; i++ {
				res += v.symbol
			}
			num = num % v.value
		} else {

		}

	}
	return res
}
