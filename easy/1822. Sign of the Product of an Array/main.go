package easy

func arraySign(nums []int) int {
	belowZero := 0

	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			belowZero++
		}
	}

	if belowZero%2 == 0 {
		return 1
	}

	return -1
}
