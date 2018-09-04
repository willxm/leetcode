package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1Head, p2Head := head, head
	for p2Head.Next != nil {

		p1Head = p1Head.Next

		if p2Head.Next.Next == nil {
			break
		}
		p2Head = p2Head.Next.Next
	}
	return p1Head
}
