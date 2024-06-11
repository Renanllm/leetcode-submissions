package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if len(stack) > 0 && isCloseParentheses(c) && isMatchParentheses(stack[len(stack)-1], c) {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, c)
	}
	return len(stack) == 0
}

func isCloseParentheses(c rune) bool {
	return c == 41 || c == 93 || c == 125
}

func isMatchParentheses(o rune, c rune) bool {
	m := map[rune]rune{
		41:  40,
		93:  91,
		125: 123,
	}
	val, ok := m[c]
	if !ok {
		return false
	}
	return val == o
}
