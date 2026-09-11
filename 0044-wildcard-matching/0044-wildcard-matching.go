func isMatch(s string, p string) bool {
	sArray := []rune(s)
	pArray := []rune(p)

	sCursor := 0
	pCursor := 0

	starIndex := -1
	matchIndex := 0

	for sCursor < len(sArray) {
		if pCursor < len(pArray) &&
			(pArray[pCursor] == '?' || pArray[pCursor] == sArray[sCursor]) {
			sCursor++
			pCursor++
			continue
		}
		if pCursor < len(pArray) && pArray[pCursor] == '*' {
			starIndex = pCursor
			matchIndex = sCursor
			pCursor++
			continue
		}

		if starIndex != -1 {
			matchIndex++
			sCursor = matchIndex
			pCursor = starIndex + 1
			continue
		}
		return false
	}
	for pCursor < len(pArray) && pArray[pCursor] == '*' {
		pCursor++
	}
	return pCursor == len(pArray)
}