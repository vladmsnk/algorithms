package tasks

import "math"

func rob(nums []int) int {
	for i := 2; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-2]
	}

	ln := len(nums)
	return int(math.Max(float64(nums[ln-2]), float64(nums[ln-1])))
}
