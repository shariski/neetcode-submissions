func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range(numbers) {
		if _, ok := seen[target-v]; ok {
			return []int{seen[target-v]+1, i+1}
		}
		seen[v] = i
	}
	return nil
}
