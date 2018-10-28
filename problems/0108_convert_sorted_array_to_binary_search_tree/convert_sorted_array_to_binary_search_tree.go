package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bst(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2

	root := &TreeNode{Val: nums[mid]}

	root.Left = bst(nums, start, mid-1)

	root.Right = bst(nums, mid+1, end)

	return root

}

func sortedArrayToBST(nums []int) *TreeNode {
	return bst(nums, 0, len(nums)-1)
}
