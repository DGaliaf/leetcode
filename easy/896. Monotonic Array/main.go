package easy

func isMonotonic(nums []int) bool {
	isDescending := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isDescending = false
			break
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && isDescending || nums[i] < nums[i+1] && !isDescending {
			return false
		}
	}

	return true
}
