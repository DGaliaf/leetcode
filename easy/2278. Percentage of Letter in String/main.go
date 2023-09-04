package easy

func percentageLetter(s string, letter byte) int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}

	for key, val := range count {
		if key == rune(letter) {
			return (val * 100) / len(s)
		}
	}

	return 0
}
