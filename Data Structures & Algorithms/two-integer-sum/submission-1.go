func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, len(nums))
	for i, v := range(nums) {
		if val, ok := seen[target - v]; ok {
			return []int{val, i}
		}
		seen[v] = i
	}
	return nil
}
