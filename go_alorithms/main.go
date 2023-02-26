package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type QueueNode struct {
	value       int
	availableAt int
}

func leastInterval(tasks []byte, n int) int {
	var counts [26]int
	queue := make([]QueueNode, 0)
	time := 0

	h := &IntHeap{}
	heap.Init(h)
	// count each entry of task
	for _, t := range tasks {
		counts[t-'A']++
	}
	//create heap
	for _, c := range counts {
		if c != 0 {
			heap.Push(h, -c)
		}
	}

	// O(n * m) m = idle
	for h.Len() > 0 || len(queue) > 0 {
		currLen := h.Len()

		time++
		if currLen != 0 {
			minInHeap := heap.Pop(h).(int)

			if minInHeap != -1 {
				queue = append(queue, QueueNode{value: minInHeap + 1, availableAt: time + n})
			}
		}

		if len(queue) > 0 {

			top := queue[0]
			if top.availableAt == time {
				queue = queue[1:]
				heap.Push(h, top.value)
			}
		}
	}

	return time
}

func main() {
	res := leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)
	fmt.Println(res)

}

//time++
//
//if h.Len() != 0 {
//	min := h.Pop().(int)
//
//	if min+1 != 0 {
//		queue = append(queue, QueueNode{value: min + 1, availableAt: time + n})
//	}
//}
//
//if len(queue) != 0 {
//	top := queue[0]
//	if top.availableAt == time {
//		queue = queue[1:]
//		h.Push(top.value)
//	}
//}
