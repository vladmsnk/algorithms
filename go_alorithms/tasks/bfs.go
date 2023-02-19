package tasks

import "container/list"

/*
golang queue implementation

queue := list.New()

append:
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

delete:
	front:=queue.Front()
    queue.Remove(front)
*/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 { //while queue is not empty
		qLen := queue.Len()
		var levelRes []int
		for i := 0; i < qLen; i++ {
			//remove node from queue
			front := queue.Front()
			enqueued := queue.Remove(front).(*TreeNode)
			if enqueued != nil {
				levelRes = append(levelRes, enqueued.Val)
				queue.PushBack(enqueued.Left)
				queue.PushBack(enqueued.Right)
			}
		}
		if len(levelRes) != 0 {
			res = append(res, levelRes)
		}
	}
	return res
}
