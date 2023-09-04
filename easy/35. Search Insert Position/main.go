package easy

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	mid := 0
	for start <= end {
		mid = (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if mid == end {
		return mid + 1
	}

	return mid
}
