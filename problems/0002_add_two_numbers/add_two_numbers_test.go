package problems

import (
	"reflect"
	"testing"
)

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				l1: makeListNode([]int{2, 4, 3}),
				l2: makeListNode([]int{5, 6, 4}),
			},
			want: makeListNode([]int{7, 0, 8}),
		},
		{
			name: "test2",
			args: args{
				l1: makeListNode([]int{1, 8}),
				l2: makeListNode([]int{0}),
			},
			want: makeListNode([]int{1, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
