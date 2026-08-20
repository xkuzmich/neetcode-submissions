func removeElement(nums []int, val int) int {
    start := 0
	end := len(nums) - 1
	count := 0
	for start <= end {
		for end > 0 && end > start && nums[end] == val {
			end--
			count++
		}
		if nums[start] == val {
			nums[start], nums[end] = nums[end], nums[start]
			end--
			count++
		}
		start++
	}
	return len(nums) - count 
}
