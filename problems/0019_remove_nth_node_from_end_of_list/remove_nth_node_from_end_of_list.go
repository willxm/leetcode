package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h1, h2 := head, head
	for i := 0; i < n; i++ {
		h1 = h1.Next
	}
	if h1 == nil {
		return head.Next
	}
	for h1.Next != nil {
		h1 = h1.Next
		h2 = h2.Next
	}
	h2.Next = h2.Next.Next
	return head
}
