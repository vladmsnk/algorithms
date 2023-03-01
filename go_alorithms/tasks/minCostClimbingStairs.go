package tasks

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	ln := len(cost)
	cost = append(cost, 0)
	// [10, 15, 20, 0]
	for i := ln - 2; i != -1; i-- {
		cost[i] = int(math.Min(float64(cost[i]+cost[i+1]), float64(cost[i]+cost[i+2])))
	}
	fmt.Println(cost)
	return 0
}
