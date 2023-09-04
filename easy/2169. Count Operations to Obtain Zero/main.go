package easy

func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}

	if num1 == num2 {
		return 1
	}

	count := 0
	for (num1 != 0 || num2 != 0) && num1 != num2 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		count++
	}
	count++

	return count
}
