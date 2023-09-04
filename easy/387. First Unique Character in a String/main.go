package easy

func firstUniqChar(s string) int {
	indexes := make(map[rune]int)
	counts := make([]int, len(s))

	for i, r := range s {
		if _, ok := indexes[r]; !ok {
			indexes[r] = i
		}

		counts[indexes[r]]++
	}

	for i := 0; i < len(counts); i++ {
		if counts[i] == 1 {
			return i
		}
	}

	return -1
}
