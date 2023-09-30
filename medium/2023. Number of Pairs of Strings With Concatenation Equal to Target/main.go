package medium

func numOfPairs(nums []string, target string) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		temp := nums[i]

		for j := 0; j < len(nums); j++ {
			if i != j {
				temp += nums[j]

				if temp == target {
					count++
				}
			}

			temp = nums[i]
		}
	}

	return count
}
