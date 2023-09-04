package easy

func kidsWithCandies(candies []int, extraCandies int) []bool {
	answer := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		cur := candies[i] + extraCandies
		isMax := true

		for j := 0; j < len(candies); j++ {
			if cur < candies[j] && i != j {
				isMax = false
			}
		}

		if isMax {
			answer[i] = true
		}
	}

	return answer
}
