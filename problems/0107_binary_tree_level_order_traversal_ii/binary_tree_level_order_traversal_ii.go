package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	bottomRes := [][]int{}

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)

	l := len(res)

	for i := 0; i < l; i++ {
		bottomRes = append(bottomRes, res[l-i-1])
	}
	return bottomRes
}
