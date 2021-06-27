// 2021.06.27

func isValid(s string) bool {
	if 1 == len(s) || 0 != len(s)%2 || ")" == string(s[0]) || "}" == string(s[0]) || "]" == string(s[0]) {
		return false
	}

	var stack []string
	for _, ch := range s {
		c := string(ch)
		if "(" == c || "{" == c || "[" == c {
			stack = append(stack, c)
		} else {
			switch c {
			case ")":
				if 0 == len(stack) || stack[len(stack)-1] != "(" {
					return false
				}
				stack = stack[:len(stack)-1]
			case "}":
				if 0 == len(stack) || stack[len(stack)-1] != "{" {
					return false
				}
				stack = stack[:len(stack)-1]
			case "]":
				if 0 == len(stack) || stack[len(stack)-1] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	return 0 == len(stack)
}
