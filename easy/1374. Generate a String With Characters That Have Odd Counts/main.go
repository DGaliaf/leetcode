package easy

func generateTheString(n int) string {
	var result string
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			result += "a"
		}

		result += "b"
	} else {
		for i := 0; i < n; i++ {
			result += "a"
		}
	}

	return result
}
