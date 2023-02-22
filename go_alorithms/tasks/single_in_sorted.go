package tasks

// длина всегда нечетная
// [1,1,3,3,4,4,8,8,9]
// [1,1,2,3,3,4,4,8,8]
// находясь в середине, нужно найти пару. Слева и справа от сердины мб четная и нечетная длина
// если длина четная, то мы ищем пару и где мы найдем пару, с той стороны нужный элемент
// если длина нчетная, то нужной элемент с противоположной стороны от пары

// [1,1,2,2,3,3,4]
// [1,2,2,3,3,4,4]

// [1, 1, 2, 2, 3, 3, 4, 5, 5]
// [1, 2, 2, 3, 3, 4, 4, 5, 5]

func singleNonDuplicate(nums []int) int {
	r, l := len(nums)-1, 0

	mid := (r + 1 - l) / 2

	for r != l {
		if nums[mid] == nums[mid-1] {
			if (mid-l)%2 == 1 {
				l = mid + 1
			} else {
				r = mid
			}

		} else if nums[mid] == nums[mid+1] {
			if (r-mid)%2 == 1 {
				r = mid - 1
			} else {
				l = mid
			}
		} else {
			return nums[mid]
		}
		mid = l + (r+1-l)/2
	}
	return nums[l]
}
