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

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				head: makeListNode([]int{1, 2, 2, 3}),
			},
			want: makeListNode([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
