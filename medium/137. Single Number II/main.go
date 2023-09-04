package medium

import "sort"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	if nums[1] != nums[0] {
		return nums[0]
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] && nums[i] != nums[i-1] {
			return nums[i]
		}
	}

	return -1
}
