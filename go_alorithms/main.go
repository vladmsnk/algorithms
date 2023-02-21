package main

import (
	"fmt"
)

// [1,0,0,0,1,1,1,0,0,0,0] k = 2

// intervals = [[8,18],[1,4],[7,10], [2,6]]
// output = [[1, 6], [7, 18]]
// [[1 4] [7 18] [2 20]]

func main() {
	//test1 := [][]int{{8, 18}, {1, 4}, {7, 10}, {5, 20}}
	//test2 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	test3 := [][]int{} //
	merged2 := merge(test3)
	fmt.Println(merged2)

}
