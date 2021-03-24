package problems

func findMin(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	} else if size == 2 {
		return min(nums[0], nums[1])
	}

	start := 0
	stop := size - 1

	for start < stop-1 {
		if nums[start] < nums[stop] {
			return nums[start]
		}

		mid := start + (stop-start)/2
		if nums[start] < nums[mid] {
			start = mid
		} else if nums[start] > nums[mid] {
			stop = mid
		} else {
			start++
		}
	}
	return min(nums[start], nums[stop])
}

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}
