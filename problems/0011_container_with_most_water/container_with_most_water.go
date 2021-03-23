package problem

func maxArea(height []int) int {
	maxArea := 0
	// for i := 0; i < len(height); i++ {
	// 	for j := i + 1; j < len(height); j++ {
	// 		area := (j - i) * minInt(height[i], height[j])
	// 		if area > maxArea {
	// 			maxArea = area
	// 		}
	// 	}
	// }
	i := 0
	j := len(height) - 1
	for i < j {
		area := (j - i) * minInt(height[i], height[j])
		if area > maxArea {
			maxArea = area
		}
		// if height[i+1]-height[i] > height[j-1]-height[j] {
		// 	i++
		// } else if height[i+1]-height[i] < height[j-1]-height[j] {
		// 	j--
		// } else {
		// 	if height[i] > height[j] {
		// 		j--
		// 	} else {
		// 		i++
		// 	}
		// }

		if height[i] > height[j] {
			j--
		} else {
			i++
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
