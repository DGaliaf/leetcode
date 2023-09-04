package easy

func detectCapitalUse(word string) bool {
	upperCount, lowerCount := 0, 0
	for _, w := range word {
		if w <= 'Z' {
			upperCount++
		} else {
			lowerCount++
		}
	}

	if upperCount == len(word) || lowerCount == len(word) {
		return true
	}

	if word[0] <= 'Z' && lowerCount == len(word)-1 && upperCount == 1 {
		return true
	}

	return false
}
