package problem

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		s    string
		cost []int
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
		// 		s:    "aabaa",
		// 		cost: []int{1, 2, 3, 4, 1},
		// 	},
		// 	want: 2,
		// },
		{
			name: "case2",
			args: args{
				s:    "aaabbbabbbb",
				cost: []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1},
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.s, tt.args.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
