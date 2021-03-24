package problems

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: -1,
		},

		{
			name: "case2",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
				target: 14,
			},
			want: 13,
		},

		{
			name: "case3",
			args: args{
				nums:   []int{2, 2},
				target: 2,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},

		{
			name: "case2",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},

		{
			name: "case3",
			args: args{
				nums:   []int{2, 2},
				target: 2,
			},
			want: []int{0, 1},
		},

		{
			name: "case4",
			args: args{
				nums:   []int{1, 3},
				target: 1,
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
