package easy

func dominantIndex(nums []int) int {
	max1, max2 := -1, -1
	maxIndex := -1

	for i, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
			maxIndex = i
		}

		if num > max2 && num < max1 {
			max2 = num
		}
	}

	if max1 >= max2*2 {
		return maxIndex
	}

	return -1
}
