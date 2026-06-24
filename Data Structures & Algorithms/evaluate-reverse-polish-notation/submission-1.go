func evalRPN(tokens []string) int {
	stack := []int{}
	for _, c := range tokens {
		if c == "+" {
			res := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], res)
		} else if c == "-" {
			res := stack[len(stack)-2] - stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], res)
		} else if c == "*" {
			res := stack[len(stack)-2] * stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], res)
		} else if c == "/" {
			res := stack[len(stack)-2] / stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], res)
		} else {
			i, _ := strconv.Atoi(c)
			stack = append(stack, i)
		}
	}
	return stack[0]
}
