package problem

func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * minInt(height[i], height[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
