package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pReversedHead *ListNode
	pNode := head
	var pPrev *ListNode

	for pNode != nil {

		pNext := pNode.Next
		if pNext == nil {
			pReversedHead = pNode
		}

		pNode.Next = pPrev

		pPrev = pNode

		pNode = pNext

	}
	return pReversedHead
}
