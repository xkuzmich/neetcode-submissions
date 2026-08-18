func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
		return strs[0]
	}
	var result []byte
	for i := 0; i < len(strs[0]); i++ {
		letter := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return string(result)
			}
			if letter != strs[j][i] {
				return string(result)
			}
		}
		result = append(result, letter)
	}
	return string(result )
}
