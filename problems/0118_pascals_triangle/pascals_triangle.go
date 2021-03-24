package problems

//A[k][n] = A[k-1][n-1] + A[k-1][n]
func generate(numRows int) [][]int {
	res := buildMatrix(numRows, numRows)

	for i := 0; i < numRows; i++ {
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func buildMatrix(m, n int) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	return res
}
