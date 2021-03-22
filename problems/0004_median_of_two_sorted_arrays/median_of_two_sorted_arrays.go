package problems

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ms := mergeSortedArrays(nums1, nums2)
	if len(ms)%2 == 1 {
		return float64(ms[len(ms)/2])
	} else {
		return (float64(ms[len(ms)/2-1]) + float64(ms[len(ms)/2])) / 2
	}
}

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	var res []int

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i != len(nums1) {
		res = append(res, nums1[i:]...)
	}

	if j != len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}
