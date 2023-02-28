package tasks

import (
	"container/heap"
	"strings"
)

type Node struct {
	symbol      byte
	cnt         int
	isAvailable int
}

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].cnt < h[j].cnt }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reorganizeString(s string) string {
	var countMap [26]int
	var res strings.Builder
	var queue []Node
	currLen := 0

	h := &NodeHeap{}
	heap.Init(h)
	for _, symbol := range s {
		countMap[symbol-'a']++
	}

	for symbol, cnt := range countMap {
		if cnt != 0 {
			q := Node{symbol: byte(symbol) + 'a', cnt: -cnt}
			heap.Push(h, q)
		}
	}

	for h.Len() != 0 {

		if h.Len() > 0 {
			heapNode := heap.Pop(h).(Node)
			res.WriteByte(heapNode.symbol)

			currLen++
			heapNode.isAvailable = currLen + 1
			heapNode.cnt++
			if heapNode.cnt != 0 {
				queue = append(queue, heapNode)
			}
		}

		if len(queue) > 0 && queue[0].isAvailable == currLen {
			top := queue[0]
			queue = queue[1:]
			heap.Push(h, top)
		}
	}

	if len(queue) > 0 {
		return ""
	}

	return res.String()
}
