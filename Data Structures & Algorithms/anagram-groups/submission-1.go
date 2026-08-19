func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	m := make(map[[26]int][]string)

	for _, word := range strs {
		var counts [26]int
		for i:=0; i<len(word); i++ {
			counts[word[i]-'a']++
		}
		m[counts] = append(m[counts], word)
	}
	var result [][]string
	for _, val := range m {
		result = append(result, val)
	}

	return result
}
