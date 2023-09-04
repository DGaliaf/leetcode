package easy

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	output := ""

	if len(word1) == 1 && len(word2) == 1 {
		return fmt.Sprintf("%s%s", word1, word2)
	}

	min := len(word1)
	if len(word2) < len(word1) {
		min = len(word2)
	}

	for i := 0; i < min; i++ {
		output += fmt.Sprintf("%c%c", word1[i], word2[i])
	}

	output += fmt.Sprintf("%s%s", word1[min:], word2[min:])

	return output
}
