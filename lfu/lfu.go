package lfu

import "container/list"

type LFUCache struct {
	FrequencyMap map[int]*list.List
	KeyValueMap  map[int]*list.Element
	capacity     int
	currCapacity int
}

type Pair struct {
	Key       int
	Val       int
	Frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{FrequencyMap: make(map[int]*list.List),
		KeyValueMap: make(map[int]*list.Element), capacity: capacity}
}

func (this *LFUCache) Get(key int) int {

}

func (this *LFUCache) Put(key int, value int) {
	var currFrequency int

	item, ok := this.KeyValueMap[key]

	if !ok {
		currFrequency = 1
		itemList := this.FrequencyMap[currFrequency] //get list with frequency 1

		itemList.PushFront(Pair{Key: key, Val: value, Frequency: 1})

		this.currCapacity++ //every time we add new node, increment current capacity

		if this.currCapacity > this.capacity { //if current capacity is equal to full capacity
			//find current min frequency in frequency map

			//delete least recently used node
			//find minimum frequency
		}
	} else {

		currFrequency = item.Value.(Pair).Frequency //getCurrent frequency

		frequencyList := this.FrequencyMap[currFrequency] //get corresponding list of this frequency

		frequencyList.Remove(item) //remove item from frequency list

		currFrequency++ //increment current frequency of item

		delete(this.KeyValueMap, key) //delete old element from key-val map

		//and add new value
		this.KeyValueMap[key] = &list.Element{Value: Pair{Key: key, Val: value, Frequency: currFrequency}}

		newItem := this.KeyValueMap[key]
		//pushNew item
		this.FrequencyMap[currFrequency].PushFront(newItem) //push to the front of new list

	}
}

func main() {
}

// 6 2 5 3
