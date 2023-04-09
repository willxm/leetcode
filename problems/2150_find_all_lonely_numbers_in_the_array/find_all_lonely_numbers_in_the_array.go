package main

import "sort"

func findLonely(nums []int) []int {
	sort.Ints(nums)

	var res []int
	for k, v := range nums {
		if ((k > 0 && v-nums[k-1] > 1) || k == 0) && ((k < len(nums)-1 && nums[k+1]-v > 1) || k == len(nums)-1) {
			res = append(res, v)
		}
	}
	return res
}
