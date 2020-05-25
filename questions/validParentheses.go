package questions

func IsValid(s string) bool {
	stack := []rune{}
	bracketsMaps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if bracketsMaps[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

func CheckValidString(s string) bool {
	leftBalance := 0
	for _, Val := range s {
		if Val == '(' || Val == '*' {
			leftBalance++
		} else {
			leftBalance--
		}
		if leftBalance < 0 {
			return false
		}
	}

	if leftBalance == 0 {
		return true
	}
	rightBalance := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || s[i] == '*' {
			rightBalance++
		} else {
			rightBalance--
		}
		if rightBalance < 0 {
			return false
		}
	}
	return true
}
