package easy

func arithmeticTriplets(nums []int, diff int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if i < j && j < k && i+j-k != i && nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					count++
				}
			}
		}
	}

	return count
}
