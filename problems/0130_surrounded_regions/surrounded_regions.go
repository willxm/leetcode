package problems

//把边界的'O'变成‘1’，然后遍历，遇见‘1’恢复成‘O’，其余都变成'X'
func solve(board [][]byte) {
	w := len(board)
	h := len(board[0])

	for i := 0; i < w; i++ {
		dfs(board, i, 0)
		dfs(board, i, h-1)
	}

	for i := 1; i < h-1; i++ {
		dfs(board, 0, i)
		dfs(board, w-1, i)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {

			if board[i][j] == 1 {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) {
		if board[i][j] == 'O' {
			board[i][j] = '1'
			dfs(board, i+1, j)
			dfs(board, i-1, j)
			dfs(board, i, j+1)
			dfs(board, i, j-1)
		}
	}
}
