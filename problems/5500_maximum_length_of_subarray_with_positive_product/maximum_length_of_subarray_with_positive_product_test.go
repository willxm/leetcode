package problems

import "testing"

func Test_getMaxLen(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case1",
		// 	args: args{
		// 		nums: []int{1, -2, -3, 4},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "case2",
		// 	args: args{
		// 		nums: []int{0, 1, -2, -3, -4},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "case3",
		// 	args: args{
		// 		nums: []int{-1, 2},
		// 	},
		// 	want: 1,
		// },

		{
			name: "case4",
			args: args{
				nums: []int{70, -18, 75, -72, -69, -84, 64, -65, 0, -82, 62, 54, -63, -85, 53, -60, -59, 29, 32, 59, -54, -29, -45, 0, -10, 22, 42, -37, -16, 0, -7, -76, -34, 37, -10, 2, -59, -24, 85, 45, -81, 56, 86},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.args.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
