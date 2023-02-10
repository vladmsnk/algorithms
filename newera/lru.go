package main

type LRUCache struct {
	Storage  map[int]*Pair
	Cache    *Pair
	Capacity int
	Head     *Pair
	Back     *Pair
}

type Pair struct {
	Key  int
	Val  int
	Prev *Pair
	Next *Pair
}

func Constructor(capacity int) LRUCache {

	return LRUCache{Capacity: capacity,
		Storage: make(map[int]*Pair, capacity), Head: &Pair{}, Back: &Pair{}}
}

func (this *LRUCache) Get(key int) int {
	currPtr, ok := this.Storage[key]
	if !ok {
		return -1
	}
	//cut its current location
	currPtrPrev := currPtr.Prev
	currPtrNext := currPtr.Next
	currPtrPrev.Next = currPtrNext
	currPtrNext.Prev = currPtrPrev

	//move to the beginning
	this.Cache.Prev = currPtr
	currPtr.Next = this.Cache
	this.Cache = this.Cache.Prev

	return currPtr.Val
}

func (this *LRUCache) Put(key int, value int) {

	pairPtr := &Pair{Key: key, Val: value}

	if this.Cache == nil {
		this.Cache = pairPtr
		//nil <-H-> *Pair{} <-B-> nil
		this.Head.Next = pairPtr
		pairPtr.Prev = this.Head

		//maybe useless
		this.Back.Prev = pairPtr
		pairPtr.Next = this.Back

		this.Storage[key] = pairPtr
		return
	}

	_, ok := this.Storage[key]
	if !ok {
		this.Storage[key] = pairPtr

		//curPtr
		this.Cache.Prev = pairPtr
		pairPtr.Next = this.Cache

		//maybe useless
		pairPtr.Prev = this.Head
		this.Head.Next = pairPtr

		this.Cache = this.Cache.Prev

		//nil <-H-> *Pair{1, 3} -> *Pair{1, 2} -> *Pair{1, 1}  <-B-> nil
		if len(this.Storage) > this.Capacity {
			preLast := this.Back.Prev.Prev
			key := preLast.Next.Key

			preLast.Next = this.Back
			this.Back.Prev = preLast
			delete(this.Storage, key)
		}

	} else { // if smth with this key already exists
		// just set new value and move this to head
		currPtr := this.Storage[key]

		//delete this ptr from list
		currPtrPrev := currPtr.Prev
		currPtrNext := currPtr.Next
		currPtrPrev.Next = currPtrNext
		currPtrNext.Prev = currPtrPrev

		//replace this ptr with newvalue one
		this.Storage[key] = pairPtr

		//push this new ptr to the beginning of the list
		//this.Cache points to beginning
		this.Cache.Prev = pairPtr
		pairPtr.Next = this.Cache

		pairPtr.Prev = this.Head
		this.Head.Next = pairPtr
		this.Cache = this.Cache.Prev
	}

}

func main() {
	Constructor(3)
}


