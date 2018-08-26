package problem

func spiralOrder(matrix [][]int) []int {
	h := len(matrix)
	w := len(matrix[0])

	res := make([]int, h*w)
	if w == 0 || h == 0 {
		return res
	}
	top := 0
	bottom := h - 1
	left := 0
	right := w - 1

	index := 0

	for {
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		top++
		if left > right || top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		right--
		if left > right || top > bottom {
			break
		}

		for i := right; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if left > right || top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			res[index] = matrix[i][left]
			index++
		}
		left++
		if left > right || top > bottom {
			break
		}

	}
	return res
}
