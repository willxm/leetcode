package main

import (
	"reflect"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				left:  19,
				right: 31,
			},
			want: []int{29, 31},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestPrimes(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
