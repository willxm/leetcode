package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v, n := 0, 0

	for {
		v, n = add(l1, l2, n)
		temp.Val = v

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			break
		}

		temp.Next = &ListNode{}
		temp = temp.Next
	}

	if n > 0 {
		temp.Next = &ListNode{Val: n}
	}

	return result
}

func add(n1, n2 *ListNode, i int) (v, n int) {
	if n1 != nil && n2 != nil {
		v = n1.Val + n2.Val + i
	}
	if v > 9 {
		v = v - 10
		n = 1
	}
	return
}
