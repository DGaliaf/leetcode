package easy

import "strings"

func canConstruct(ransomNote string, magazine string) bool {
	for _, r := range ransomNote {
		i := strings.Index(magazine, string(r))

		if i == -1 {
			return false
		}

		magazine = magazine[:i] + magazine[i+1:]
	}

	return true
}
