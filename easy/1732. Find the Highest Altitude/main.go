package easy

import "math"

func largestAltitude(gain []int) int {
	attitudes := make([]int, len(gain)+1)
	attitudes[0] = 0

	max := math.MinInt

	for i := 0; i < len(gain); i++ {
		attitudes[i+1] += attitudes[i] + gain[i]

		if attitudes[i] > max {
			max = attitudes[i]
		}
	}
	if attitudes[len(attitudes)-1] > max {
		max = attitudes[len(attitudes)-1]
	}

	return max
}
