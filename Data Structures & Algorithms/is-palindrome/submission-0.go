func isPalindrome(s string) bool {
	var filtered []rune
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			filtered = append(filtered, r)
		}
	}
	st := string(filtered)
	for i := 0; i < len(st)/2; i++ {
		if st[i] != st[len(st)-i-1] {
			return false
		}
	}
	return true
}
