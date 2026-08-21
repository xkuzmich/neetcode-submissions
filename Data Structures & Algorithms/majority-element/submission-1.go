func majorityElement(nums []int) int {
    balance := 0
	candidate := nums[0]

	for _, num := range nums {
		if balance == 0 {
			candidate = num
		}
		if num == candidate {
			balance++
		} else {
			balance--
		}
	}
	return candidate
}
