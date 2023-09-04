package easy

import "math"

func findMaxAverage(nums []int, k int) float64 {
	result := float64(math.MinInt)

	for i := 0; i <= len(nums)-k; i++ {
		temp := 0.0
		for j := i; j < i+k; j++ {
			temp += float64(nums[j])
		}

		if temp > result {
			result = temp
		}
	}

	return result / float64(k)
}
