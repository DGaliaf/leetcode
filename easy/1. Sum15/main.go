package easy

func twoSum(nums []int, target int) []int {
	s := map[int][]int{}

	for i := 0; i < len(nums); i++ {
		_, ok := s[nums[i]]
		if ok {
			return []int{s[nums[i]][0], i}
		} else {
			s[target-nums[i]] = []int{i, 1}
		}
	}

	return []int{0, 0}
}
