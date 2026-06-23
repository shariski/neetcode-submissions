func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	pProduct, sProduct := 1, 1
	for i := 0; i < len(nums); i++ {
		pProduct *= nums[i]
		sProduct *= nums[len(nums)-i-1]
		prefix[i] = pProduct
		suffix[len(nums)-i-1] = sProduct
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, suffix[i+1])
		} else if i == len(nums)-1 {
			result = append(result, prefix[i-1])
		} else {
			result = append(result, prefix[i-1]*suffix[i+1])
		}
	}

	return result
}
