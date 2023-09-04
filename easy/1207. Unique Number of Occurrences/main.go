package easy

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	occured := make(map[int]bool)

	for _, n := range arr {
		count[n]++
	}

	for _, v := range count {
		if _, ok := occured[v]; ok {
			return false
		} else {
			occured[v] = true
		}
	}

	return true
}
