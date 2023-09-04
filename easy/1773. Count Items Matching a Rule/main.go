package easy

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var ruleIndex int
	switch ruleKey {
	case "type":
		ruleIndex = 0
		break
	case "color":
		ruleIndex = 1
		break
	case "name":
		ruleIndex = 2
		break
	}

	count := 0
	for i := 0; i < len(items); i++ {
		if items[i][ruleIndex] == ruleValue {
			count++
		}
	}

	return count
}
