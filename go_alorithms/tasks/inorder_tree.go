package tasks

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) != 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		//pop element from stack
		index := len(stack) - 1
		curr = stack[len(stack)-1]
		res = append(res, curr.Val)
		stack = stack[:index]

		curr = curr.Right
	}
	return res
}
