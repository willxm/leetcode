package problems

func setZeroes(matrix [][]int) {
	// find zero index
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := range rows {
		if rows[i] {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}

	for i := range cols {
		if cols[i] {
			for j := range matrix {
				matrix[j][i] = 0
			}
		}
	}
}
