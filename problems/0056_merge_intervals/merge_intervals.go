package problems

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type list []Interval

func merge(intervals []Interval) []Interval {

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(list(intervals))

	res := make([]Interval, 0, len(intervals))
	temp := intervals[0]

	for i := 1; i < len(intervals); i++ {

		if intervals[i].Start <= temp.End {
			temp.End = max(temp.End, intervals[i].End)
			continue
		}

		res = append(res, temp)
		temp = intervals[i]

	}
	res = append(res, temp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].Start < l[j].Start
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
