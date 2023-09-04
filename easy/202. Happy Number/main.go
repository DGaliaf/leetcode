package easy

func isHappy(n int) bool {
	if n == 1 || n == 7 {
		return true
	}

	if n >= 2 && n <= 9 {
		return false
	}

	mod := 0
	sum := 0
	for n >= 10 {
		mod += n % 10
		sum += (n % 10) * (n % 10)

		n /= 10
	}
	sum += n * n

	if n == 1 && mod == 0 {
		return true
	}

	return isHappy(sum)
}
