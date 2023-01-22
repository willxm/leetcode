package main

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] > 0 {
			return -1
		}
		if abs(nums[i]) == nums[j] {
			return nums[j]
		} else if abs(nums[i]) > nums[j] {
			i++
		} else {
			j--
		}
	}
	return -1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
