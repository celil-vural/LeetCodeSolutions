func isValidSudoku(board [][]byte) bool {
	for i := range 9 {
		if !isValidUnit(board[i]) {
			return false
		}
	}
	for j := range 9 {
		column := make([]byte, 9)
		for i := range 9 {
			column[i] = board[i][j]
		}
		if !isValidUnit(column) {
			return false
		}
	}
	for boxRow := range 3 {
		for boxCol := range 3 {
			subBox := make([]byte, 9)
			idx := 0
			for i := boxRow * 3; i < boxRow*3+3; i++ {
				for j := boxCol * 3; j < boxCol*3+3; j++ {
					subBox[idx] = board[i][j]
					idx++
				}
			}
			if !isValidUnit(subBox) {
				return false
			}
		}
	}
	return true
}
func isValidUnit(unit []byte) bool {
	seen := make(map[byte]bool)
	for _, num := range unit {
		if num < 48 || num > 57 {
			continue
		}
		if seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}
