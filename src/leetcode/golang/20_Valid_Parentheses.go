package golang

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := make([]string, 0)
	var str string
	for _, v := range s {
		str = string(v)
		if _, ok := m[str]; ok { // 匹配到右括号
			l := len(stack)
			if l == 0 {
				return false
			}
			pop := stack[l-1]
			stack = stack[:l-1]
			if m[str] != pop {
				return false
			}
		} else { // 匹配到左括号
			stack = append(stack, str)
		}
	}
	return len(stack) == 0
}
