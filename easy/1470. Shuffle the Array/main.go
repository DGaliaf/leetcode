package easy

func shuffle(nums []int, n int) []int {
	answer := make([]int, n*2)

	for left, right, i := 0, n, 0; right < n*2; left, right, i = left+1, right+1, i+2 {
		answer[i], answer[i+1] = nums[left], nums[right]
	}

	return answer
}
