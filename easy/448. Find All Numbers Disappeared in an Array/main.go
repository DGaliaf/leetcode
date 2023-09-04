package easy

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		correct := nums[i] - 1

		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}
