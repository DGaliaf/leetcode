package easy

func countGoodSubstrings(s string) int {
	result := 0

	for i := 0; i < len(s)-2 && len(s) >= 3; i++ {
		sub := s[i : i+3]

		if sub[0] != sub[1] && sub[0] != sub[2] && sub[1] != sub[2] {
			result++
		}
	}

	return result
}
