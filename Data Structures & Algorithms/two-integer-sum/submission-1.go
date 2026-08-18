func twoSum(nums []int, target int) []int {
    diffMap := make(map[int]int)

	for i, val := range nums {
		diff := target - val 

		if j, ok := diffMap[diff]; ok {
			return []int{j, i}
		}
		diffMap[val] = i
	}
	return []int{0,0}
}
