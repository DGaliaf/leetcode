package easy

import "math"

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	max := math.MinInt
	answer := math.MinInt
	for key, val := range counts {
		if val > max {
			answer = key
			max = val
		}
	}

	return answer
}
