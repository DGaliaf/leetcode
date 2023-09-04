package easy

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	connections := make(map[uint8]string)
	seen := make(map[string]bool)
	for i := 0; i < len(words); i++ {

		if _, ok := connections[pattern[i]]; !ok {
			if seen[words[i]] {
				return false
			}

			connections[pattern[i]] = words[i]
			seen[words[i]] = true
		}

		if words[i] != connections[pattern[i]] {
			return false
		}
	}

	return true
}
