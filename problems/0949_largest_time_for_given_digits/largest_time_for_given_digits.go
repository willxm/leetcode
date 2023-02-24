package main

import "fmt"

func largestTimeFromDigits(arr []int) string {
	ress := permute(arr)
	maxT := -1
	for _, v := range ress {
		h := v[0]*10 + v[1]
		m := v[2]*10 + v[3]
		if h < 24 && m < 60 {
			maxT = max(maxT, h*100+m)
		}
	}
	if maxT == -1 {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", maxT/100, maxT%100)
}

func permute(nums []int) [][]int {
	var ans [][]int
	var dfs func(l []int, temp []int)
	dfs = func(l []int, temp []int) {
		if len(l) == 0 {
			ans = append(ans, temp)
		}
		for i := 0; i < len(l); i++ {
			n := append([]int{}, l...)
			dfs(append(n[:i], n[i+1:]...), append(temp, l[i]))
		}
	}
	dfs(nums, []int{})
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
