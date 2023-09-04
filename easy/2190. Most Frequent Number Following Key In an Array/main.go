package easy

func mostFrequent(nums []int, key int) int {
	counts := make(map[int]int)

	res, max := 0, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			counts[nums[i+1]]++
			if counts[nums[i+1]] > max {
				res = nums[i+1]
				max = counts[nums[i+1]]
			}
		}
	}

	return res
}
