package problems

import (
	"reflect"
	"testing"
)

func Test_mergeSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
				[]int{1, 3},
				[]int{2},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortedArrays(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 3},
				[]int{2},
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				[]int{1, 2},
				[]int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
