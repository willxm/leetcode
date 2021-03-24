package problems

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {

	tNums := typeInt(nums)

	sort.Sort(tNums)

	res := ""

	first := tNums[0]

	l := len(tNums)

	if l == 1 {
		return strconv.Itoa(tNums[0])
	}

	for i := 0; i < l; i++ {

		if first == 0 {
			if i == l-1 {
				return "0"
			}
			first = tNums[i+1]
			continue
		}

		res += strconv.Itoa(tNums[i])
	}

	return res

}

type typeInt []int

func (t typeInt) Len() int {
	return len(t)
}

func (t typeInt) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t typeInt) Less(i, j int) bool {
	stri := strconv.Itoa(t[i])
	strj := strconv.Itoa(t[j])

	str1 := stri + strj

	str2 := strj + stri

	return str1 > str2
}
