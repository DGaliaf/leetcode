package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum *int, low, high int) {
	if root != nil {
		if root.Val >= low && root.Val <= high {
			*sum += root.Val
		}

		dfs(root.Left, sum, low, high)
		dfs(root.Right, sum, low, high)
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	dfs(root, &sum, low, high)

	return sum
}
