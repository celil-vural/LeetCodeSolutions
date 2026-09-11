func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		subset := append([]int(nil), nums[:i]...)
		subset = append(subset, nums[i+1:]...)
		for _, p := range permuteUnique(subset) {
			temp := append([]int{num}, p...)
			if !isExistInSlice(result, temp) {
				result = append(result, temp)
			}
		}
	}
	return result
}
func isExistInSlice(slice [][]int, value []int) bool {
	for _, v := range slice {
		exists := true
		for i := range v {
			if v[i] != value[i] {
				exists = false
				break
			}
		}
		if exists {
			return true
		}
	}
	return false
}
