package problems

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}

	p := nums[0]
	minP := nums[0]
	maxP := nums[0]

	for i := 1; i < l; i++ {

		tmp := maxP
		maxP = max(max(maxP*nums[i], nums[i]), minP*nums[i])
		minP = min(min(tmp*nums[i], nums[i]), minP*nums[i])
		p = max(maxP, p)
	}
	return p
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
