package easy

import "strings"

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	for i := 1; i <= l/2; i++ {
		if l%i == 0 && strings.Repeat(s[:i], l/i) == s {
			return true
		}
	}

	return false
}
