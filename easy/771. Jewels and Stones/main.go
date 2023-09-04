package easy

func numJewelsInStones(jewels string, stones string) int {
	jewelsDict := make(map[rune]int)
	for _, jewel := range jewels {
		jewelsDict[jewel] = 1
	}

	count := 0
	for _, stone := range stones {
		count += jewelsDict[stone]
	}

	return count
}
