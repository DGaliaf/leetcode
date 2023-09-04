package easy

func isVowel(r byte) bool {
	vowels := []byte{'a', 'e', 'o', 'u', 'i', 'A', 'E', 'O', 'U', 'I'}

	for _, vowel := range vowels {
		if r == vowel {
			return true
		}
	}

	return false
}

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1

	bs := []byte(s)

	for left < right {
		if isVowel(bs[left]) && isVowel(bs[right]) {
			bs[left], bs[right] = bs[right], bs[left]

			left++
			right--
		}

		if !isVowel(bs[left]) {
			left++
		}

		if left == right {
			break
		}

		if !isVowel(bs[right]) {
			right--
		}
	}

	return string(bs)
}
