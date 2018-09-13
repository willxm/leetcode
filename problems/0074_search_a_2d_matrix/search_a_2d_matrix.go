package problem

func searchMatrix(matrix [][]int, target int) bool {

	line := len(matrix)
	if line < 1 {
		return false
	}
	row := len(matrix[0])
	if row < 1 {
		return false
	}

	for i := 0; i < line; i++ {
		for j := row - 1; j > -1; j-- {
			if target > matrix[i][j] {
				break
			} else if target < matrix[i][j] {
				continue
			} else {
				return true
			}
		}
	}
	return false
}
