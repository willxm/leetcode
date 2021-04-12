package problems

func findDuplicate(nums []int) int {
	d := make([]byte, len(nums))
	for _, v := range nums {
		if d[v] != 1 {
			d[v] = 1
		} else {
			return v
		}
	}
	return 0
}
