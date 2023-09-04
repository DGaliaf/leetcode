package medium

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)

	res := make([][]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		groups[groupSizes[i]] = append(groups[groupSizes[i]], i)

		if len(groups[groupSizes[i]]) == groupSizes[i] {
			res = append(res, groups[groupSizes[i]])
			groups[groupSizes[i]] = nil
		}
	}

	return res
}
