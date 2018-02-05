package problems

func twoSum(nums []int, target int) []int {
	l := len(nums)
	start := 0
	end := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] < target {
				continue
			}
			if nums[i]+nums[j] > target {
				continue
			}
			if nums[i]+nums[j] == target {
				start = i
				end = j
			}
		}
	}
	return []int{start, end}
}
