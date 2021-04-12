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

func findDuplicate1(a []int) int {
	slow, fast := a[0], a[a[0]]
	for slow != fast {
		slow, fast = a[slow], a[a[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = a[slow], a[fast]
	}

	return slow
}
