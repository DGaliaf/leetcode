package easy

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func leftRigthDifference(nums []int) []int {
	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
	}

	res := make([]int, len(nums))
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		right := prefixSum - nums[i] - leftSum
		s := right - leftSum
		res[i] = abs(s)

		leftSum += nums[i]
	}

	return res
}
