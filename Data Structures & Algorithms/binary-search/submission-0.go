func search(nums []int, target int) int {
	for i, v := range(nums) {
		if target == v {
			return i
		}
	}
	return -1
}
