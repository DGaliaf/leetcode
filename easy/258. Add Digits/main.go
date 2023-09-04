package easy

func addDigits(num int) int {
	if num >= 0 && num <= 9 {
		return num
	}

	return 1 + ((num - 1) % 9)
}
