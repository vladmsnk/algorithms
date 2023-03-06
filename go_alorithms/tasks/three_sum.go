package tasks

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	mpCheck := make(map[[3]int]bool)

	for k, target := range nums {
		target = -1 * target

		i, j := 0, len(nums)-1
		for i < j {

			if nums[i]+nums[j] > target || j == k {
				j--
			} else if nums[i]+nums[j] < target || i == k {
				i++
			} else if nums[i]+nums[j] == target && i != k && j != k {

				tmp := []int{nums[i], nums[j], nums[k]}
				sort.Ints(tmp)
				a := *(*[3]int)(tmp)
				if !mpCheck[a] {
					mpCheck[a] = true
					res = append(res, tmp)
				}
				i++
				j--
			}
		}
	}
	return res
}
