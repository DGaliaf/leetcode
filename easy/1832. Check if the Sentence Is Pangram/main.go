package easy

func checkIfPangram(sentence string) bool {
	count := make(map[rune]struct{})

	for _, r := range sentence {
		if _, ok := count[r]; !ok {
			count[r] = struct{}{}
		}
	}

	return len(count) == 26
}
