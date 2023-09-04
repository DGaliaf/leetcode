package easy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minimumSum(num int) int {
	numSlice := strings.Split(fmt.Sprintf("%d", num), "")
	sort.Strings(numSlice)

	a, _ := strconv.Atoi(numSlice[0] + numSlice[len(numSlice)-1])
	b, _ := strconv.Atoi(numSlice[1] + numSlice[len(numSlice)-2])

	return a + b
}
