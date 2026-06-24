func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	maxLen := 0
	for v := range m {
		if !m[v-1] {
			l := 1
			for m[v+l] {
				l++
			}
			if l > maxLen {
				maxLen = l
			}
		}
	}
	return maxLen
}
