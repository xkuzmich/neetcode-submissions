func removeElement(nums []int, val int) int {
	end := len(nums)
	for start:=0; start < end; {
		if nums[start] == val {
			nums[start], nums[end-1] = nums[end-1], nums[start]
			end--
		} else {
		start++
		}
	}
	return end
}