package easy

func moveZeroes(nums []int) {
	indexToPlace := 0

	for i, _ := range nums {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = 0
			nums[indexToPlace] = temp

			indexToPlace++
		}
	}
}
