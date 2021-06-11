package valid

func isValid(s string) bool {
	// we're going to use a stack ( slice tricks) and a map
	stack := []rune{}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		// switch case, if opening parentheses, add to stack, if closing parentheses,
		//pull from stack and should match map
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			// check if stack is empty first OR if the top of stack doesnt match map
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			// else we're good, pull from stack
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
