package easy

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxProfit(prices []int) int {
	currentMin := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-currentMin)
		currentMin = min(currentMin, prices[i])
	}

	return maxProfit
}
