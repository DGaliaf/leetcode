package easy

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	count1 := make(map[int]bool)
	for _, n := range nums1 {
		count1[n-1] = true
	}

	count2 := make(map[int]bool)
	for _, n := range nums2 {
		count2[n-1] = true
	}

	count3 := make(map[int]bool)
	for _, n := range nums3 {
		count3[n-1] = true
	}

	res := make([]int, 0)
	for i := 0; i < 100; i++ {
		if (count1[i] && count2[i]) || (count1[i] && count3[i]) || (count3[i] && count2[i]) {
			res = append(res, i+1)
		}
	}

	return res
}
