package problems

func sortColors(nums []int) {
	n0, n1, n2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			n0++
		} else if nums[i] == 1 {
			n1++
		} else if nums[i] == 2 {
			n2++
		}
	}
	for i := 0; i < n0; i++ {
		nums[i] = 0
	}

	for i := 0; i < n1; i++ {
		nums[n0+i] = 1
	}

	for i := 0; i < n2; i++ {
		nums[n0+n1+i] = 2
	}

}
