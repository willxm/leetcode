package problems

func getMaxLen(nums []int) int {
	// sort.Ints(nums)

	maxLen := 0

	for i := 0; i < len(nums); i++ {
		tmpMaxLen := 0

		tmp := nums[i]
		if tmp > 0 {
			tmp = 1
		} else if tmp < 0 {
			tmp = -1
		}

		if tmp > 0 && maxLen == 0 {
			maxLen = 1
		}

		if nums[i] == 0 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {

			if nums[j] == 0 {
				i = j
				break
			}

			if nums[j] > 0 {
				nums[j] = 1
			} else if nums[j] < 0 {
				nums[j] = -1
			}

			tmp = tmp * nums[j]

			tmpMaxLen = j - i + 1

			// if tmpMaxLen == 13 {
			// 	tmpMaxLen = 13
			// }

			// if tmp > 0 {
			// 	tmpMaxLen = i - j + 1
			// }

			if tmp > 0 {
				// tmp = 1
				if tmpMaxLen > maxLen {
					maxLen = tmpMaxLen
				}
			}
		}
	}
	return maxLen
}
