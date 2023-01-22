package main

import "testing"

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, -1, -2, -3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxK(tt.args.nums); got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}
