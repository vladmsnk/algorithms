package tasks

import (
	"sort"
)

// intervals = [[8,18],[1,4],[7,10], [2,6]]
// output = [[1, 6], [7, 18]]

func merge(intervals [][]int) [][]int {
	var merged [][]int

	compareFunc := func(i, j int) bool { return intervals[i][0] < intervals[j][0] }
	sort.Slice(intervals, compareFunc)
	for i, j := 0, 1; j < len(intervals)+1; i, j = i+1, j+1 {
		if j != len(intervals) && intervals[j][1] >= intervals[i][1] && intervals[i][1] >= intervals[j][0] {
			intervals[j][0] = intervals[i][0]
		} else if j != len(intervals) && intervals[j][1] <= intervals[i][1] {
			intervals[j][0] = intervals[i][0]
			intervals[j][1] = intervals[i][1]
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
