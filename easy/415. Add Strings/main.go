package easy

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	if len(num2) < len(num1) {
		num1, num2 = num2, num1
	}
	dif := len(num2) - len(num1)

	result := make([]string, 0)
	carry := 0
	for i := len(num1) - 1; i >= 0; i-- {
		sum := int(num1[i]-'0') + int(num2[i+dif]-'0') + carry
		carry = sum / 10

		result = append(result, strconv.Itoa(sum%10))
	}

	for i := dif - 1; i >= 0; i-- {
		sum := int(num2[i]-'0') + carry
		carry = sum / 10

		result = append(result, strconv.Itoa(sum%10))
	}

	if carry == 1 {
		result = append(result, "1")
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return strings.Join(result, "")
}
