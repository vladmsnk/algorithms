package tasks

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	//if rotated len(nums) times
	if nums[l] <= nums[r] {
		return nums[l] //min is the first element
	}
	mid := 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid]
		}
		if nums[mid+1] <= nums[r] { //right side is sorted
			r = mid
		} else if nums[l] <= nums[mid] {
			l = mid + 1
		}
	}
	return -1
}

// 8, 9, 1, 2, 3, 4, 6, 7
