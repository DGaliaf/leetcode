package easy

func maximum69Number(num int) int {
	var digits []int

	n := 10
	for num > 0 {
		digits = append(digits, num%n)
		num /= n
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 6 {
			digits[i] = 9
			break
		}
	}

	var res int
	for i := len(digits) - 1; i >= 0; i-- {
		res += digits[i]
		res *= 10
	}

	return res / 10
}
