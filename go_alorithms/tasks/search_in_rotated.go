package tasks

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := 0

	for l < r {
		mid = l + (r-l)/2

		//simple binary search
		if nums[l] <= nums[r] {
			if target > nums[mid] {
				l = mid + 1
			} else if target < nums[mid] {
				r = mid
			} else {
				return mid
			}
		} else if nums[mid+1] <= nums[r] {
			//target might be in sorted right
			if target >= nums[mid+1] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}

		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}

	mid = l + (r-l)/2
	if nums[mid] == target {
		return mid
	}

	return -1
}
