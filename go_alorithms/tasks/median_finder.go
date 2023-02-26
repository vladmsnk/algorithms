package tasks

import "container/heap"

type MedianFinder struct {
	MaxHeap *IntHeap
	MinHeap *IntHeap
}

func Constructor() MedianFinder {
	m := MedianFinder{MaxHeap: &IntHeap{},
		MinHeap: &IntHeap{}}
	heap.Init(m.MaxHeap)
	heap.Init(m.MinHeap)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.MaxHeap, num)

	if this.MaxHeap.Len()-this.MinHeap.Len() > 1 || this.MinHeap.Len() > 0 && (*this.MaxHeap)[0] < (*this.MinHeap)[0] {
		maxValue := heap.Pop(this.MaxHeap)
		heap.Push(this.MinHeap, maxValue)
	}

	if this.MinHeap.Len()-this.MaxHeap.Len() > 1 {
		minValue := heap.Pop(this.MinHeap)
		heap.Push(this.MaxHeap, minValue)
	}

}

//
//func (this *MedianFinder) FindMedian() float64 {
//	if this.MinHeap.Len() > this.MaxHeap.Len() {
//
//	}
//}
