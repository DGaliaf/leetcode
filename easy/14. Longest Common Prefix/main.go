package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstWord := strs[0]

	for i := 0; i < len(firstWord); i++ {
		for _, word := range strs[1:] {
			if i == len(word) || firstWord[i] != word[i] {
				return firstWord[:i]
			}
		}
	}

	return firstWord
}
