package easy

func containsDuplicate(nums []int) bool {
	counts := map[int]struct{}{}

	for _, num := range nums {
		_, ok := counts[num]

		if !ok {
			counts[num] = struct{}{}
		} else {
			return true
		}
	}

	return false
}
