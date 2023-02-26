package tasks

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
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
	for _, c := range counts {
		if c != 0 {
			h.Push(-1 * c)
		}
	}

	// O(n * m) m = idle
	for h.Len() > 0 || len(queue) > 0 {
		currLen := h.Len()

		time++
		if currLen != 0 {
			minInHeap := heap.Remove(h, 0).(int)

			heap.Fix(h, h.Len()-1)
			if minInHeap != -1 {
				// zeros wont go to queue
				queue = append(queue, QueueNode{value: minInHeap + 1, availableAt: time + n})
			}
		}

		if len(queue) > 0 && queue[0].availableAt == time {

			top := queue[0]

			queue = queue[1:]
			h.Push(top.value)
		}
	}

	return time
}
