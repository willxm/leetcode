package problem

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	size := len(logs)

	res := make([]int, n)
	stack := make([]int, 0, size)
	endPoint := 0
	lastEndPoint := 0

	for _, log := range logs {
		fid, soe, ts := parse(log)

		if soe == "start" {
			endPoint = ts
			if len(stack) > 0 {
				res[stack[len(stack)-1]] += endPoint - lastEndPoint
			}
			stack = append(stack, fid)
			lastEndPoint = endPoint
		} else {
			endPoint = ts + 1
			res[fid] += endPoint - lastEndPoint
			stack = stack[:len(stack)-1]
			lastEndPoint = endPoint
		}
	}

	return res
}

func parse(log string) (int, string, int) {
	ss := strings.Split(log, ":")
	funcID, _ := strconv.Atoi(ss[0])
	timeStamp, _ := strconv.Atoi(ss[2])
	return funcID, ss[1], timeStamp
}
