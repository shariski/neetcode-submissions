func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}
		m[s[right]] = right
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
