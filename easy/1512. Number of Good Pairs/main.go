package easy

func numIdenticalPairs(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	res := 0
	for _, v := range counts {
		if v > 1 {
			res += v * (v - 1) / 2
		}
	}

	return res
}
