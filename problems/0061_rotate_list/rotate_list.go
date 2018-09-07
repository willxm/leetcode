package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	count := 1
	pNode := head
	for pNode.Next != nil {
		pNode = pNode.Next
		count++
	}
	k = k % count
	if k == 0 {
		return head
	}
	pNode.Next = head
	for i := 0; i < count-k; i++ {
		pNode = pNode.Next
	}
	nex := pNode.Next
	// pNode = prev.Next
	pNode.Next = nil
	return nex
}
