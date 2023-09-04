package easy

func maximumWealth(accounts [][]int) int {
	result := 0

	for _, account := range accounts {
		temp := 0
		for _, balance := range account {
			temp += balance
		}

		if temp > result {
			result = temp
		}
	}

	return result
}
