package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{[]int{1, 2, 3, 1}, 3}, want: true},
		{name: "2", args: args{[]int{1, 0, 1, 1}, 1}, want: true},
		{name: "3", args: args{[]int{1, 2, 3, 1, 2, 3}, 2}, want: false},
		{name: "4", args: args{[]int{99, 99}, 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
