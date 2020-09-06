package problem

import "testing"

func Test_numTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
				nums1: []int{7, 4},
				nums2: []int{5, 2, 8, 9},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				nums1: []int{7, 7, 8, 3},
				nums2: []int{1, 2, 9, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTriplets(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("numTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
