package easy

import "math"

func getMaximumGenerated(n int) int {
	result := math.MinInt
	nums := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			nums[i] = i
		}

		doubled := 2 * i
		if doubled >= 2 && doubled <= n {
			nums[doubled] = nums[i]
		}

		if doubled+1 >= 2 && doubled+1 <= n {
			nums[doubled+1] = nums[i] + nums[i+1]
		}

		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}
