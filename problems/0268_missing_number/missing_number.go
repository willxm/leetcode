package problem

func missingNumber(nums []int) int {
	x := 0
	l := len(nums)
	for i := 0; i <= l; i++ {
		x ^= i
	}
	for i := 0; i < l; i++ {
		x ^= nums[i]
	}
	return x
}
