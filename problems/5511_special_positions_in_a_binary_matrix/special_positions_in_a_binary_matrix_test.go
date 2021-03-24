package problems

import "testing"

func Test_numSpecial(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					[]int{1, 0, 0},
					[]int{0, 0, 1},
					[]int{1, 0, 0},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecial(tt.args.mat); got != tt.want {
				t.Errorf("numSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
