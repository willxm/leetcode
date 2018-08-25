package problem

// [1,2,3],
// [4,5,6],
// [7,8,9]

// [7,4,1],
// [8,5,2],
// [9,6,3]

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			// fmt.Println(m[i][j], "<-", m[n-j-1][i])
			m[i][j] = m[n-j-1][i]
			// fmt.Println(m[n-j-1][i], "<-", m[n-i-1][n-j-1])
			m[n-j-1][i] = m[n-i-1][n-j-1]
			// fmt.Println(m[n-i-1][n-j-1], "<-", m[j][n-i-1])
			m[n-i-1][n-j-1] = m[j][n-i-1]
			// fmt.Println(m[j][n-i-1], "<-", temp)
			m[j][n-i-1] = temp
		}
	}
}
