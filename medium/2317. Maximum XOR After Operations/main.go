package medium

func maximumXOR(nums []int) (max int) {
	for _, num := range nums {
		max |= num
	}

	return
}
