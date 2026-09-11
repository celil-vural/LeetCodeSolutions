func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	jumps := 0
	current := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == current {
			current = farthest
			jumps++
		}
	}
	return jumps
}
