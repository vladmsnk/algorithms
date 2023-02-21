package tasks

import "container/list"

func invertTree(root *TreeNode) *TreeNode {
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
				//swith nodes
				tmp := enqueued.Left
				enqueued.Left = enqueued.Right
				enqueued.Right = tmp

				levelRes = append(levelRes, enqueued.Val)
				queue.PushBack(enqueued.Left)
				queue.PushBack(enqueued.Right)
			}
		}
	}
	return root
}
