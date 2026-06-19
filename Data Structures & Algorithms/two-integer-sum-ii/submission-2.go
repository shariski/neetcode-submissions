func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for i := 0; i < len(numbers); i++ {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left+1, right+1}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return nil
}
