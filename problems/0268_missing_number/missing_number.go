package problem

func missingNumber(nums []int) int {
	x := 0
	for i := 0; i <= len(nums); i++ {
		x ^= i
	}
	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}
	return x
}
