func firstMissingPositive(nums []int) int {
	slices.Sort(nums)
	missing, zeroCursor := 0, 0
	for i := range nums {
		if nums[i] <= 0 {
			zeroCursor++
			continue
		}
		if i == zeroCursor && nums[i] != 1 {
			return 1
		}
		if i > zeroCursor && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] != missing+1 {
			return missing + 1
		}
		missing = nums[i]
	}
	res := nums[len(nums)-1]
	if res < 0 {
		res = 0
	}
	return res + 1
}
