package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	pHead := head

	for pHead.Next != nil {

		if pHead.Val == pHead.Next.Val {
			pHead.Next = pHead.Next.Next
		} else {
			pHead = pHead.Next
		}
	}
	return head
}
