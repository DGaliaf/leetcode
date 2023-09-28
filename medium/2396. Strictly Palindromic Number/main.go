package medium

import "strconv"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(val string) bool {
	return val[:len(val)/2] == reverseString(val[len(val)/2:])
}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-1; i++ {
		nInBase := strconv.FormatInt(int64(n), i)

		if !isPalindrome(nInBase) {
			return false
		}
	}

	return true
}
