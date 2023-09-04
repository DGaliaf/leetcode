package easy

import "strings"

func capitalizeTitle(title string) string {
	words := make([]string, 0, len(title))
	words = strings.Split(title, " ")

	var res string
	for _, word := range words {
		if len(word) <= 2 {
			res += strings.ToLower(word)
		} else {
			res += strings.Title(strings.ToLower(word))
		}

		res += " "
	}

	return strings.TrimSpace(res)
}
