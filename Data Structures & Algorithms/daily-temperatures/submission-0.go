func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[j] = i-j
		}
		stack = append(stack, i)
	}
	return result
}
