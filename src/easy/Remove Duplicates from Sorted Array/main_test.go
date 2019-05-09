package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{[]int{0, 0, 1, 1, 2, 3}}, want: 4},
		// {name: "2", args: args{[]int{1, 1, 2}}, want: 2},
		// {name: "3", args: args{[]int{9, 9, 1, 1, 2, 3}}, want: 4},
		{name: "4", args: args{[]int{0, 1, 2, 3, 3, 4}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
