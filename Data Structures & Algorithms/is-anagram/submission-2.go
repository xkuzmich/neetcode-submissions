func isAnagram(s string, t string) bool {
	checkMap := make(map[byte]int)

	if len(s) != len(t) { return false}

	for i, _ := range s {
		checkMap[s[i]] += 1
		checkMap[t[i]] -= 1
	}

	for _, val := range checkMap {
		if val != 0 {
			return false
		}
	}

	return true
 }
