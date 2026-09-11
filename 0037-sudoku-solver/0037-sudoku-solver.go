func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 46 {
				continue
			}
			var num byte
			for num = 49; num <= 57; num++ {
				if !isValid(board, row, col, num) {
					continue
				}
				board[row][col] = num
				if solve(board) {
					return true
				}
				board[row][col] = 46
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for c := 0; c < 9; c++ {
		if board[row][c] == num {
			return false
		}
	}
	for r := 0; r < 9; r++ {
		if board[r][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}
	return true
}
