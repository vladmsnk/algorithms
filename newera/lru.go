package main

import "fmt"

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
	//Move to the beginning
	if len(this.Storage) > 1 && currPtr != this.Cache {

		prevNode := currPtr.Prev
		nextNode := currPtr.Next

		prevNode.Next = nextNode
		nextNode.Prev = prevNode

		currPtr.Next = this.Cache
		currPtr.Prev = this.Head

		this.Cache.Prev = currPtr
		this.Head.Next = currPtr

		this.Cache = currPtr
	}

	return currPtr.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	pairPtr := &Pair{Key: key, Val: value}

	if this.Cache == nil {
		this.Cache = pairPtr
		//H <- *Pair{} -> B
		this.Cache.Next = this.Back
		this.Cache.Prev = this.Head
		//H <-> *Pair{} <-> B
		this.Head.Next = this.Cache
		this.Back.Prev = this.Cache
		// store pointer in map
		this.Storage[key] = this.Cache
		return
	}

	_, ok := this.Storage[key]
	if !ok {
		// add to the beginning
		pairPtr.Next = this.Cache
		pairPtr.Prev = this.Head

		this.Cache.Prev = pairPtr
		this.Head.Next = pairPtr

		this.Cache = pairPtr

		this.Storage[key] = pairPtr

		if len(this.Storage) > this.Capacity {

			preLast := this.Back.Prev.Prev
			keyToDelete := preLast.Next.Key

			preLast.Next = this.Back
			this.Back.Prev = preLast
			delete(this.Storage, keyToDelete)
		}

	} else { // if smth with this key already exists
		// just set new value and move this to head

		currPtr := this.Storage[key]
		currPtr.Val = value

		if len(this.Storage) > 1 && currPtr != this.Cache {
			//swap
			prevNode := currPtr.Prev
			nextNode := currPtr.Next

			prevNode.Next = nextNode
			nextNode.Prev = prevNode

			currPtr.Next = this.Cache
			currPtr.Prev = this.Head
			this.Cache.Prev = currPtr
			this.Head.Next = currPtr

			this.Cache = currPtr
		}
	}
}
func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 4)

	lru.Put(2, 2)
	lru.Put(5, 5)
	lru.Put(6, 6)
	lru.Put(7, 7)
	lru.Put(8, 8)

	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)

	val := lru.Get(8)
	val = lru.Get(8)
	fmt.Println(val)

	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)

	//FirstNode:  3  :  2
	//SecondNode:  3  :  2
	//

}

//func main() {
//	lru := Constructor(3)
//	lru.Put(1, 1)
//	lru.Put(1, 2)
//	fmt.Println("____Test1, replace value by key____")
//	//lru.Put(1, 2)
//	fmt.Println("List: ", lru.Cache.Key, " : ", lru.Cache.Val)
//
//	currKey := lru.Cache.Key
//	fmt.Println("Map: ", lru.Storage[currKey].Key, " : ", lru.Storage[currKey].Val)
//
//	fmt.Println("____Test2, add several nodes____")
//	//Nodes that were added should be in the beginning
//	//*Pair{3, 4} should be at the first place and *Pair{2, 3} at the second
//	lru.Put(2, 3)
//	lru.Put(3, 4)
//	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
//	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
//	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)
//
//	fmt.Println("____Test3, add node so that the capacity will be stepped over ____")
//	//Node *Pair{4, 6} should be in the first place.
//	//Node *Pair{1, 2}  should be deleted
//	lru.Put(4, 6)
//	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
//	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
//	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)
//
//	fmt.Println("____Test4, add node so that the capacity will be stepped over again____")
//	// First node is Pair{5, 7}
//	// Second node is Pair{4, 6}
//	// Third node is Pair{3, 4}
//	lru.Put(5, 7)
//	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
//	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
//	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)
//
//	fmt.Println("____Test5, get existing key____")
//	val := lru.Get(3)
//	//val should be 4
//	fmt.Println("val :", val)
//	fmt.Println("FirstNode: ", lru.Cache.Key, " : ", lru.Cache.Val)
//	fmt.Println("SecondNode: ", lru.Cache.Next.Key, " : ", lru.Cache.Next.Val)
//	fmt.Println("ThirdNode: ", lru.Cache.Next.Next.Key, " : ", lru.Cache.Next.Next.Val)
//
//}
