package easy

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if s == "" {
		return true
	}

	left := 0
	for right := 0; right < len(t); right++ {
		if s[left] == t[right] {
			left++

			if left >= len(s) {
				break
			}
		}
	}

	return left >= len(s)
}
