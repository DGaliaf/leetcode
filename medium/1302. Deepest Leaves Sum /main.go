package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, level int, maxLevel *int, levels map[int][]int, sum int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		levels[level] = append(levels[level], root.Val)
	}

	if level > *maxLevel {
		*maxLevel = level
	}

	dfs(root.Left, level+1, maxLevel, levels, sum)
	dfs(root.Right, level+1, maxLevel, levels, sum)
}

func deepestLeavesSum(root *TreeNode) int {
	sum := 0
	maxLevel, levels := 0, make(map[int][]int)

	dfs(root, 0, &maxLevel, levels, sum)

	for _, v := range levels[maxLevel] {
		sum += v
	}

	return sum
}
