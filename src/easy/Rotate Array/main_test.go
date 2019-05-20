package main

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}},
		{name: "2", args: args{[]int{-1, -100, 3, 99}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}
