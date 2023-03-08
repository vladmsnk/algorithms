package tasks

func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
		} else {
			top := stack[len(s)-1]
			stack = stack[1 : len(s)-1]
			if r != top {
				return false
			}
		}
	}
	return true
}
