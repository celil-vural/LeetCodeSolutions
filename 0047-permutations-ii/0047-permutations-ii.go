func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	current := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			result = append(result, append([]int(nil), current...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			current = append(current, nums[i])
			backtrack()
			current = current[:len(current)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}
