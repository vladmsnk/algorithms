package tasks

import "math"

// [1,1,1,0,0,0,1,1,1,1,0]

func longestOnes(nums []int, k int) int {
	currK := k
	maxOnes := 0

	for l, r := 0, 0; r < len(nums); {
		if nums[r] == 1 || currK > 0 {
			if nums[r] == 0 {
				currK--
			}
			r++
			maxOnes = int(math.Max(float64(maxOnes), float64(r-l)))
		} else {
			if nums[l] == 0 {
				currK++
			}
			l++
		}
	}
	return maxOnes
}
