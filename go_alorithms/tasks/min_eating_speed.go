package tasks

func findMax(nums []int) int {
	max := 0

	for _, r := range nums {
		if r > max {
			max = r
		}
	}
	return max
}

func countHours(p, k int) int {
	cur := p / k
	if p%k != 0 {
		return cur + 1
	}
	return cur
}

func ableToEat(piles []int, h, k int) bool {
	i := 0

	for h > 0 {
		hours := countHours(piles[i], k)
		h -= hours
		i++
	}
	if h < 0 {
		return false
	}
	return true
}
func minEatingSpeed(piles []int, h int) int {
	max := findMax(piles)

	if h == len(piles) {
		return max
	}
	l, r := 0, max

	k := 0
	for l < r {
		k = l + (r-l)/2
		able := ableToEat(piles, h, k)
		if able {
			r = k
		} else {
			l = k + 1
		}
	}
	return k
}

// [3,6,7,11]   h = 8 k = 5
