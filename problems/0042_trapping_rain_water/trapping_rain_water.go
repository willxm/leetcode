package main

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]

	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])
		right[length-1-i] = max(right[length-i], height[length-1-i])
	}

	water := 0
	for i := 0; i < length; i++ {
		water += min(left[i], right[i]) - height[i]
	}

	return water
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
