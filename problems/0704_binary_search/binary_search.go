package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	i, j := 0, len(nums)-1
	for {
		if i > j {
			return -1
		}
		mid := (i + j) / 2
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return mid
		}
	}
}
