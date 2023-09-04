package easy

func repeatedCharacter(s string) byte {
	seen := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := seen[r]; ok {
			return byte(r)
		}

		seen[r] = struct{}{}
	}

	return byte(0)
}
