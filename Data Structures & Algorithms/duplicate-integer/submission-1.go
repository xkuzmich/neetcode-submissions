func hasDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }

    cache := make(map[int]string)

    for _, elem := range nums {
        _, exist := cache[elem]
        if exist {
            return true
        }
        cache[elem] = ""
    }
    return false
}
