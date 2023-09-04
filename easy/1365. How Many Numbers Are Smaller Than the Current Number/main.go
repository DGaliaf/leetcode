package easy

func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		smaller := 0
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i] > nums[j] {
					smaller++
				}
			}
		}
		output[i] = smaller
	}

	return output
}
