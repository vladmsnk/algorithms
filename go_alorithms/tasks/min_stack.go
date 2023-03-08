package tasks

type MinStack struct {
	head *StackNode
	len  int
}

type StackNode struct {
	val     int
	currMin int
	next    *StackNode
}

func ConstructorStack() MinStack {
	return MinStack{head: &StackNode{}}
}

func (this *MinStack) Push(val int) {
	currMin := 0
	if this.len == 0 {
		currMin = val
	} else {
		currMin = min(this.head.currMin, val)
	}
	newNode := &StackNode{val: val, currMin: currMin}
	newNode.next = this.head
	this.head = newNode
	this.len++
}

func (this *MinStack) Pop() {
	if this.len == 0 {
		return
	}
	this.head = this.head.next
	this.len--
}

func (this *MinStack) Top() int {

	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.currMin
}

// stack: 6 7 5 4

// mins: 6 6 5 4

//when pop: if stack[len(stack) - 1] == mins[len(mins) - 1] mins=
