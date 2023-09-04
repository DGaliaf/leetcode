package easy

func sum(nums []int) int {
	s := 0

	for _, num := range nums {
		s += num
	}

	return s
}

func missingNumber(nums []int) int {
	n := len(nums)

	return n*(n+1)/2 - sum(nums)
}
