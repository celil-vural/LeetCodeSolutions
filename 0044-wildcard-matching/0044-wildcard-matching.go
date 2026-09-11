func isMatch(s string, p string) bool {
	sArray := []rune(s)
	pArray := []rune(p)
	// 0 = unknown
	// 1 = true
	// 2 = false
	memo := make(map[[2]int]uint8)

	var match func(sCursor, pCursor int) bool

	match = func(sCursor, pCursor int) bool {
		key := [2]int{sCursor, pCursor}

		if result, ok := memo[key]; ok {
			return result == 1
		}
		if pCursor == len(pArray) {
			result := sCursor == len(sArray)
			if result {
				memo[key] = 1
			} else {
				memo[key] = 2
			}
			return result
		}
		var result bool
		if pArray[pCursor] == '*' {
			result = match(sCursor, pCursor+1)
			if !result && sCursor < len(sArray) {
				result = match(sCursor+1, pCursor)
			}
		} else if sCursor < len(sArray) &&
			(pArray[pCursor] == '?' || sArray[sCursor] == pArray[pCursor]) {
			result = match(sCursor+1, pCursor+1)
		}
		if result {
			memo[key] = 1
		} else {
			memo[key] = 2
		}
		return result
	}
	return match(0, 0)
}