package validparenthesis

func isValid(s string) bool {

	var stack []string

	for _, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			stack = append(stack, string(v))
		} else if string(v) == ")" || string(v) == "}" || string(v) == "]" {
			if len(stack) > 0 && ((string(v) == ")" && stack[len(stack)-1] == "(") || (string(v) == "}" && stack[len(stack)-1] == "{") || (string(v) == "]" && stack[len(stack)-1] == "[")) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
