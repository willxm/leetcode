package problem

func numSpecial(mat [][]int) int {
	r := len(mat)
	c := len(mat[0])

	count := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 1 {
				if sum(mat, i, j) == 1 {
					count++
				}
			}
		}
	}
	return count
}

func sum(mat [][]int, xr, xc int) int {
	r := len(mat)
	c := len(mat[0])

	s := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == xr {
				s += mat[xr][j]
			}
			if j == xc {
				s += mat[i][xc]
			}
		}
	}
	return s - 1
}
