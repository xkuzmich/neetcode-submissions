func reverseString(s []byte) {
	end := len(s) - 1
	for start :=0; start < end; {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
