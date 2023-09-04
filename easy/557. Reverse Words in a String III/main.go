package easy

import "strings"

func reverseWords(s string) string {
	splitedString := strings.Split(s, " ")

	res := ""
	for _, str := range splitedString {
		bStr := []byte(str)

		for i := 0; i < len(str)/2; i++ {
			temp := bStr[i]
			bStr[i] = bStr[len(str)-i-1]
			bStr[len(str)-i-1] = temp
		}

		res += string(bStr)
		res += " "
	}

	return res[:len(res)-1]
}
