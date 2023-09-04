package easy

import "strings"

func isBanned(word string, banned []string) bool {
	for _, banWord := range banned {
		if word == banWord || word == "" {
			return true
		}
	}

	return false
}

func isSymbol(r byte) bool {
	if r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == '\'' {
		return true
	}

	return false
}

func mostCommonWord(paragraph string, banned []string) string {
	count := make(map[string]int)

	temp := ""
	for i := 0; i < len(paragraph); i++ {
		if isSymbol(paragraph[i]) {
			temp = strings.ToLower(temp)

			if !isBanned(temp, banned) {
				count[temp]++
			}

			temp = ""
		} else {
			temp += string(paragraph[i])
		}
	}

	if temp != "" {
		count[strings.ToLower(temp)]++
	}

	max := 0
	res := ""
	for key, val := range count {
		if val > max {
			max = val
			res = key
		}
	}

	return res
}
