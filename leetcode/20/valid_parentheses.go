package main

func isValid(s string) bool {
	var stack []uint8
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				switch s[i] {
				case ')':
					if stack[len(stack)-1] != '(' {
						return false
					}
					stack = stack[:len(stack)-1]
				case ']':
					if stack[len(stack)-1] != '[' {
						return false
					}
					stack = stack[:len(stack)-1]
				case '}':
					if stack[len(stack)-1] != '{' {
						return false
					}
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
