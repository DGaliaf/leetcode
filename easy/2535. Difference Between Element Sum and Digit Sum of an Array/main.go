package easy

import "strconv"

func abs(a, b int) int {
	c := a - b

	if c < 0 {
		return -c
	}

	return c
}

func differenceOfSum(nums []int) int {
	elementSum, digitSum := 0, 0

	for _, num := range nums {
		elementSum += num

		for _, digit := range strconv.Itoa(num) {
			d, _ := strconv.Atoi(string(digit))

			digitSum += d
		}
	}

	return abs(elementSum, digitSum)
}
