package problems

func permute(nums []int) [][]int {
	var res [][]int
	start := 0
	end := len(nums) - 1

	permuteLocal(&res, nums, start, end)

	return res
}

func permuteLocal(res *[][]int, nums []int, start, end int) {
	if start == end {
		var tmp []int
		for i := 0; i <= end; i++ {
			tmp = append(tmp, nums[i])
		}
		*res = append(*res, tmp)
	} else {
		for i := start; i <= end; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			permuteLocal(res, nums, start+1, end)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}
