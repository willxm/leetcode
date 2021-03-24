package problems

func searchRange(nums []int, target int) []int {
	index := binarySearch(nums, target)

	if index == -1 {
		return []int{-1, -1}
	}

	left, right := index, index

	for {

		// if left-1 >= 0 && nums[left-1] == target {
		// 	left--
		// }

		// if right+1 <= len(nums)-1 && nums[right+1] == target {
		// 	right++
		// }

		// if left-1 < 0 && right+1 > len(nums)-1 {
		// 	break
		// }

		// if (left-1 >= 0 && nums[left-1] != target) && (right+1 <= len(nums)-1 && nums[right+1] != target) {
		// 	break
		// }

		if left-1 >= 0 {
			if nums[left-1] == target {
				left--
			} else {
				break
			}
		} else {
			break
		}
	}

	for {
		if right+1 <= len(nums)-1 {
			if nums[right+1] == target {
				right++
			} else {
				break
			}
		} else {
			break
		}
	}

	return []int{left, right}
}

func binarySearch(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
	}

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
