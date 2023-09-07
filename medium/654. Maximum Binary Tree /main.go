package medium

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxIndex(nums []int) int {
	max, maxIndex := math.MinInt, 0

	for i, n := range nums {
		if n > max {
			max = n
			maxIndex = i
		}
	}

	return maxIndex
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	idx := maxIndex(nums)
	return &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
}
