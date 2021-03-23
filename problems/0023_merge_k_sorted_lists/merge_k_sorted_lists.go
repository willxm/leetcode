package problems

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var mergeHead *ListNode

	if l1.Val < l2.Val {
		mergeHead = l1
		mergeHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		mergeHead = l2
		mergeHead.Next = mergeTwoLists(l1, l2.Next)
	}
	return mergeHead
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	base := lists[0]
	for i := 1; i < len(lists); i++ {
		base = mergeTwoLists(base, lists[i])
	}
	return base
}
