package main

import "testing"

func Test_largestTimeFromDigits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: "23:41",
		},
		{
			name: "test2",
			args: args{
				arr: []int{5, 5, 5, 5},
			},
			want: "",
		},
		{
			name: "test3",
			args: args{
				arr: []int{0, 0, 0, 0},
			},
			want: "00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTimeFromDigits(tt.args.arr); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
