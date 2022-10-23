package problems

import "sort"

type Cc struct {
	Char  rune
	Count int64
}

func frequencySort(s string) string {
	cm := make(map[rune]int64)
	for _, r := range s {
		cm[r]++
	}
	var ccs []Cc
	for r, c := range cm {
		ccs = append(ccs, Cc{
			Char:  r,
			Count: c,
		})
	}
	sort.Slice(ccs, func(i, j int) bool { return ccs[i].Count > ccs[j].Count })
	var res []rune
	for _, v := range ccs {
		for i := 0; i < int(v.Count); i++ {
			res = append(res, v.Char)
		}
	}
	return string(res)
}
