package easy

import "math"

func average(salary []int) float64 {
	min := math.MaxInt
	max := math.MinInt

	sum := 0.0
	for i := 0; i < len(salary); i++ {
		if salary[i] > max {
			max = salary[i]
		}

		if salary[i] < min {
			min = salary[i]
		}

		sum += float64(salary[i])
	}

	return (sum - (float64(max) + float64(min))) / float64(len(salary)-2)
}
