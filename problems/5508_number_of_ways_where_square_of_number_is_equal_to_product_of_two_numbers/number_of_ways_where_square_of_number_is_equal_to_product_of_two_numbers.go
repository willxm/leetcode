package problem

func numTriplets(nums1 []int, nums2 []int) int {
	var res int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			for k := j; k < len(nums2); k++ {
				if nums1[i]*nums1[i] == nums2[j]*nums2[k] {
					res++
				}
			}
		}
	}
	return res
}
