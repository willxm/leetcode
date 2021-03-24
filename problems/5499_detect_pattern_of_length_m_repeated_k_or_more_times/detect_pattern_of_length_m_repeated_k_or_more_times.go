package problems

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i < len(arr)-m; i++ {
		tmpK := 1
		var tmpPattern []int

		tmpPattern = arr[i : i+m]

		for j := i + m; j < len(arr); j = j + m {
			if j+m > len(arr) {
				break
			}

			if arrEqual(tmpPattern, arr[j:j+m]) {
				tmpK++
			} else {
				break
			}

		}
		if tmpK == k {
			return true
		}
	}
	return false
}

func arrEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
