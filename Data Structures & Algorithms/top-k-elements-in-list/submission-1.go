func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	bucket := make([][]int, len(nums)+1)
	for i, sum := range freq {
		bucket[sum] = append(bucket[sum], i)
	}

	result := []int{}
	for i := len(nums); i >= 0 && len(result) < k; i-- {
		result = append(result, bucket[i]...)
	}
	return result
}
