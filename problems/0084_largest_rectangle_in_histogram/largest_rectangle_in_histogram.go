package problem

func largestRectangleArea(heights []int) int {
	var maxArea int
	// var minHeight int

	// maxHeight := heights[0]

	if len(heights) == 1 {
		return heights[0]
	}

	for i := 0; i < len(heights); i++ {
		tmpMax := heights[i] * 1
		tmpMinHeight := heights[i]

		if i != 0 {
			if heights[i-1] > heights[i] {
				continue
			}
		}

		for j := i + 1; j < len(heights); j++ {

			if heights[j] < tmpMinHeight {
				tmpMinHeight = heights[j]
			}

			if tmpMax < tmpMinHeight*(j-i+1) {
				tmpMax = tmpMinHeight * (j - i + 1)
			}
		}
		if maxArea < tmpMax {
			maxArea = tmpMax
		}
	}
	return maxArea
}
