package tasks

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, tmp *ListNode

	for head.Next != nil {
		prev = head
		head = head.Next
		prev.Next = tmp
		tmp = prev
	}
	return head
}
