func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    cache := make(map[byte]int)

    for i := range len(s) {
        cache[s[i]]++
        cache[t[i]]--
    }

    for _, v := range cache {
        if v != 0 {
            return false
        }
    }
    return true

}
