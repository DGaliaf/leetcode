package easy

import "strings"

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	return len(strings.Fields(s))
}
