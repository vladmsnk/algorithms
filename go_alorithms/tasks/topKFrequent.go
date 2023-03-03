package tasks

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {

	cntMap := make(map[int]int)
	for _, i := range nums {
		cntMap[i]++
	}
	frequency := make(map[int][]int)

	for key, val := range cntMap {
		frequency[val] = append(frequency[val], key)
	}
	var keys []int

	for freq := range frequency {
		keys = append(keys, freq)
	}
	sort.Ints(keys)
	var res []int
	cnt := 0

	for i := len(keys) - 1; cnt != k; i-- {
		currKey := keys[i]
		currValue := frequency[currKey]
		if len(currValue) <= k-cnt {
			res = append(res, currValue...)
			cnt += len(currValue)
		} else {
			res = append(res, currValue[k-cnt])
			cnt = k
		}

	}

	return res
}
