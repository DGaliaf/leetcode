package easy

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	counter := 0

	sNum := fmt.Sprintf("%d", num)
	for i := 0; i < len(sNum)-(k-1); i++ {
		temp := sNum[i : i+k]

		tNum, _ := strconv.Atoi(temp)

		if tNum != 0 && num%tNum == 0 && len(temp) == k {
			counter++
		}
	}

	return counter
}
