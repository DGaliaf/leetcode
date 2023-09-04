package easy

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	var temp []int

	prefixSum := 0
	for _, num := range nums {
		prefixSum += num
		temp = append(temp, prefixSum)
	}

	return NumArray{
		nums: temp,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}

	return this.nums[right] - this.nums[left-1]
}
