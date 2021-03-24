package problems

import "testing"

func Test_containsPattern(t *testing.T) {
	type args struct {
		arr []int
		m   int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case1",
		// 	args: args{
		// 		arr: []int{1, 2, 4, 4, 4, 4},
		// 		m:   1,
		// 		k:   3,
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "case2",
		// 	args: args{
		// 		arr: []int{1, 2, 1, 2, 1, 3},
		// 		m:   2,
		// 		k:   3,
		// 	},
		// 	want: false,
		// },
		// {
		// 	name: "case3",
		// 	args: args{
		// 		arr: []int{2, 2},
		// 		m:   1,
		// 		k:   2,
		// 	},
		// 	want: true,
		// },
		{
			name: "case4",
			args: args{
				arr: []int{2, 2, 1, 2, 2, 1, 1, 1, 2, 1},
				m:   2,
				k:   2,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsPattern(tt.args.arr, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("containsPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
