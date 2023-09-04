package easy

func pivotIndex(nums []int) int {
	result := -1

	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	cur := 0
	for i := 0; i < len(nums); i++ {
		temp := total - cur - nums[i]
		if cur == temp {
			result = i
			break
		}

		cur += nums[i]
	}

	return result
}
