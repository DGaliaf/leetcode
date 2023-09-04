package medium

func sortColors(nums []int) {
	zeroIndex := 0
	twoIndex := len(nums) - 1

	for i := 0; i <= twoIndex; i++ {
		if nums[i] == 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		} else if nums[i] == 2 {
			nums[i], nums[twoIndex] = nums[twoIndex], nums[i]
			twoIndex--
			i--
		}
	}
}
