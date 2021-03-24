package problems

import "testing"

func Test_modifyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case1",
		// 	args: args{
		// 		s: "?zs",
		// 	},
		// 	want: "azs",
		// },
		{
			name: "case2",
			args: args{
				s: "wz???a??n",
			},
			want: "wzacbacan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyString(tt.args.s); got != tt.want {
				t.Errorf("modifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
