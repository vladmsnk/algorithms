package tasks

func findMaxConsecutiveOnes(nums []int) int {
	maxLen := 0

	for l, r := 0, 1; r < len(nums)+1; r++ {

		if nums[l] == 0 {
			l++
		}
		if (r == len(nums) || nums[r] == 0) && nums[r-1] == 1 {
			if r-l >= maxLen {
				maxLen = r - l
			}
			l = r
		}
	}
	return maxLen
}
