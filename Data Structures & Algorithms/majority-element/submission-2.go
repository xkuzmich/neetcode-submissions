func majorityElement(nums []int) int {
    balance := 0
	var candidate int

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
