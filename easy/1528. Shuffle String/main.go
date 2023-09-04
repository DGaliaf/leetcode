package easy

import "strings"

func restoreString(s string, indices []int) string {
	answer := make([]string, len(s))

	for i := 0; i < len(indices); i++ {
		answer[indices[i]] = string(s[i])
	}

	return strings.Join(answer, "")
}
