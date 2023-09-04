package easy

func isPalindrome(x int) bool {
	copyX := x
	revX := 0

	r := 0
	for copyX > 0 {
		r = copyX % 10
		revX = revX*10 + r

		copyX /= 10
	}

	if x == revX {
		return true
	}

	return false
}
