package easy

import (
	"strconv"
	"strings"
)

func hammingWeight(num uint32) int {
	binaryRepresentation := strconv.FormatUint(uint64(num), 2)

	return strings.Count(binaryRepresentation, "1")
}
