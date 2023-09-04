package medium

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	output := ""
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" {
			continue
		}

		output += words[i]
		output += " "
	}

	return output[:len(output)-1]
}
