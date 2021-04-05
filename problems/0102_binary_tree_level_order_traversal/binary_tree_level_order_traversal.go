package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

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
	return res
}

// FIXME: print
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}

	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	tmp := []int{}

	for len(queue) != 0 {

		qr := queue[0]
		queue = queue[1:]
		tmp = append(tmp, qr.Val)

		res = append(res, tmp)
		tmp = []int{}

		if qr.Left != nil {
			queue = append(queue, qr.Left)
		}

		if qr.Right != nil {
			queue = append(queue, qr.Right)
		}

	}

	return res

}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var quene []*TreeNode
	quene = append(quene, root)
	var res [][]int
	var currentRes []int
	for len(quene) > 0 {
		currentLen := len(quene)
		NewQuene := []*TreeNode{}
		currentRes = []int{}
		for i := 0; i < currentLen; i++ {

			node := quene[i]
			currentRes = append(currentRes, node.Val)
			if node.Left != nil {
				NewQuene = append(NewQuene, node.Left)
			}
			if node.Right != nil {
				NewQuene = append(NewQuene, node.Right)
			}
		}
		quene = NewQuene
		res = append(res, currentRes)
	}
	return res

}
