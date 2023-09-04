package easy

import "strings"

func sortSentence(s string) string {
	words := strings.Split(s, " ")

	res := make([]string, len(words)+1)
	for _, word := range words {
		pos := word[len(word)-1] - '0'

		res[pos] = word[:len(word)-1]
	}

	return strings.Join(res, " ")[1:]
}
