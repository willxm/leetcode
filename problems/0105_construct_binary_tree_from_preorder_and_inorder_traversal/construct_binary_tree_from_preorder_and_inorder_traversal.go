package problems

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	lp := len(preorder)
	li := len(inorder)
	if lp != li || lp < 1 || li < 1 {
		return nil
	}
	root := constructCore(preorder, 0, lp-1, inorder, 0, li-1)
	return root
}

func constructCore(preOrder []int, startPre, endPre int, inOrder []int, startIn, endIn int) *TreeNode {

	rootValue := preOrder[startPre]

	node := &TreeNode{
		Val: rootValue,
	}

	if startPre == endPre {
		return node
	}

	index := 0
	for index = startIn; index < endIn; index++ {
		if inOrder[index] == rootValue {
			break
		}
	}

	leftLen := index - startIn
	rightLen := endIn - index

	if leftLen > 0 {
		node.Left = constructCore(preOrder, startPre+1, startPre+leftLen, inOrder, startIn, index-1)
	}
	if rightLen > 0 {
		node.Right = constructCore(preOrder, endPre-rightLen+1, endPre, inOrder, index+1, endIn)
	}
	return node

}
