package problems

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{1, 2},
				[]int{3},
				[]int{1, 3},
				[]int{2, 3},
				[]int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
