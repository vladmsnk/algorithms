package tasks

func hasPathSum(root *TreeNode, targetSum int) bool {
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) != 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		//pop element from stack
		index := len(stack) - 1
		curr = stack[index]
		stack = stack[:index]

		curr = curr.Right
	}
}
