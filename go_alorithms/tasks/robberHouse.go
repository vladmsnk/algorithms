package tasks

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	nums[1] = Max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		nums[i] = Max(nums[i-1], nums[i]+nums[i-2])
	}

	return nums[len(nums)-1]
}
