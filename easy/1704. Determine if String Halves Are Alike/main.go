package easy

func isVowel(r byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, v := range vowels {
		if v == r {
			return true
		}
	}

	return false
}

func halvesAreAlike(s string) bool {
	count1, count2 := 0, 0
	for i, j := 0, len(s)/2; j < len(s); i, j = i+1, j+1 {
		if isVowel(s[i]) {
			count1++
		}

		if isVowel(s[j]) {
			count2++
		}
	}

	return count1 == count2
}
