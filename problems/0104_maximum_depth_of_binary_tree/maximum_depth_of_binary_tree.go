package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
