func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]uint8][]string)

	for _, word := range strs {
		var counts [26]uint8
		for i:=0; i<len(word); i++ {
			counts[word[i]-'a']++
		}
		groups[counts] = append(groups[counts], word)
	}
	result := make([][]string, 0, len(groups))
	for _, val := range groups {
		result = append(result, val)
	}

	return result
}
