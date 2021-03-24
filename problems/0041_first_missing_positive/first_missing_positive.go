package problems

func firstMissingPositive(nums []int) int {
	nm := make(map[int]int)
	min := 999999
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < min && nums[i] > 0 {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		nm[nums[i]] = 1
	}

	if min != 1 {
		return 1
	}

	for j := min; j <= max+1; j++ {
		if _, ok := nm[j]; !ok {
			return j
		}
	}
	return 1
}
