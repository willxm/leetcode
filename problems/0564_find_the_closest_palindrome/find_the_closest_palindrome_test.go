package problems

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				n: "1234",
			},
			want: "1221",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
