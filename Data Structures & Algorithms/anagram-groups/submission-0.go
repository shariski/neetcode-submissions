func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	result := [][]string{}
	for _, str := range(strs) {
		var key [26]byte
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		m[key] = append(m[key], str)
	}

	for _, group := range(m) {
		result = append(result, group)
	}

	return result
}
