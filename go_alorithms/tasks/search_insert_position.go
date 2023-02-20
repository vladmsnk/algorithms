package tasks

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.
*/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := (r - l + 1) / 2

	for l != r {
		if nums[mid] <= target {
			l = mid
			mid = mid + (r-l+1)/2
		} else {
			r = mid - 1
			mid = (r - l + 1) / 2
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
