package easy

func removeElement(nums []int, val int) int {
	pos := 0

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			count--
			pos++
		} else if pos != 0 {
			nums[i-pos] = nums[i]
			nums[i] = val
		}
		count++
	}

	return count
}
