package easy

func plusOne(digits []int) []int {
	end := len(digits) - 1
	digits[end]++

	r := digits[end] / 10
	if r != 0 {
		digits[end] = 0

		for j := end - 1; j >= 0; j-- {
			digits[j] += r
			r = digits[j] / 10

			if r == 0 {
				break
			} else {
				digits[j] = 0
			}
		}
	}

	if r != 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
