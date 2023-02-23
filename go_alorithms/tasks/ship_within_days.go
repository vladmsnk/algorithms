package tasks

func isPossibleToShip(weights []int, days, capacity int) bool {
	cnt := 0
	partialSum := 0

	for _, el := range weights {
		if el > capacity {
			return false
		}
		if partialSum+el > capacity {
			partialSum = 0
			cnt++
		}
		partialSum += el
	}
	if cnt < days {
		return true
	}
	return false
}

func sum(slice []int) (res int) {
	for _, el := range slice {
		res += el
	}
	return res
}

func shipWithinDays(weights []int, days int) int {
	lBound, rBound := 0, sum(weights)

	capacity := lBound + (rBound-lBound)/2
	for lBound != rBound {

		couldBeShipped := isPossibleToShip(weights, days, capacity)
		if couldBeShipped {
			rBound = capacity
		} else {
			lBound = capacity + 1
		}
		capacity = lBound + (rBound-lBound)/2
	}
	return capacity
}
