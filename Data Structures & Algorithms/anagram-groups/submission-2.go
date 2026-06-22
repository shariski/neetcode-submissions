func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string, len(strs))
	
	for _, str := range strs {
		var key [26]byte
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		m[key] = append(m[key], str)
	}

	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}
