func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = next(res)
	}
	return res
}
func next(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); {
		count := 1
		for i+count < len(s) && s[i] == s[i+count] {
			count++
		}

		result.WriteString(strconv.Itoa(count))
		result.WriteByte(s[i])
		i += count
	}
	return result.String()
}