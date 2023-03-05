package tasks

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, i := range nums {
		set[i] = false
	}

	maxLen := 0

	for _, i := range nums {

		_, ok := set[i-1]
		curLen := 1
		if !ok {
			j := i + 1
			_, ok := set[j]
			for ok {
				curLen++
				_, ok = set[j+1]
				j++
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}

	}

	return maxLen
}
