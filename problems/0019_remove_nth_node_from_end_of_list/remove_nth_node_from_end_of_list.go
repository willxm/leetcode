package problem

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := getEndNode(head, n)
	if node == nil {
		return head.Next
	}
	node.Next = node.Next.Next
	return head
}

func getEndNode(head *ListNode, n int) *ListNode {
	node := head

	for head != nil {
		if n < 0 {
			node = node.Next
		}
		n--
		head = head.Next
	}
	if n == 0 {
		return nil
	}
	return node
}
