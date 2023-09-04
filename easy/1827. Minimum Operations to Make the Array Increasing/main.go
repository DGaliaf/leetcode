package easy

func minOperations(nums []int) int {
	operations := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			operations += (nums[i] - nums[i+1]) + 1
			nums[i+1] = nums[i] + 1
		}
	}

	return operations
}
