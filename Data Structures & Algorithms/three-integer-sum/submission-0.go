func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}
	for i := 0; i < n-2; i++ {
		// impossible to get sum = 0 if num > 0 because array already sorted
		// in other words, just calculate elements until 0
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L, R := i+1, n-1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum < 0 {
				L++
			} else if sum > 0 {
				R--
			} else {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			}
		}
	}
	return result
}
