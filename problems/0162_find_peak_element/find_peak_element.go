package problems

func findPeakElement(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return 0
	} else if size == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}
	for i := 1; i < size; i++ {
		if i == size-1 && nums[i] > nums[i-1] {
			return size - 1
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}
