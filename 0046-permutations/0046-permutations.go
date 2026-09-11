func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var result [][]int
	for i, num := range nums {
		subset := append([]int(nil), nums[:i]...)
		subset = append(subset, nums[i+1:]...)
		for _, p := range permute(subset) {
			result = append(result, append([]int{num}, p...))
		}
	}
	return result
}
