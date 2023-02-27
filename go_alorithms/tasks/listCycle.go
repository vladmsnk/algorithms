package tasks

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	// 1, 2, 3
	frst := head
	scnd := head.Next.Next
	for frst.Next != nil && scnd != nil {
		if frst == scnd {
			return true
		}
		frst = frst.Next
		if scnd.Next != nil {
			scnd = scnd.Next.Next
		} else {
			return false
		}
	}
	return false
}
