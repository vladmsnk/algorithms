package tasks

import (
	"fmt"
	"math"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	res := make([]string, 0)
	bgn := 0
	curr := 0

	for i, j := 0, 1; j < len(nums)+1; i, j = i+1, j+1 {
		prev := nums[i]
		if j != len(nums) {
			curr = nums[j]
		}

		if math.Abs(float64(curr-prev)) != 1 || j == len(nums) {
			if bgn == i {
				res = append(res, strconv.Itoa(nums[bgn]))
			} else {
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[bgn]), strconv.Itoa(prev)))
			}
			bgn = j
		}
	}
	return res
}
