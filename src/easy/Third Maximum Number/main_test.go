package main

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{3, 2, 1}}, want: 1},
		{name: "2", args: args{[]int{2, 1}}, want: 2},
		{name: "3", args: args{[]int{2, 2, 3, 1}}, want: 1},
		{name: "4", args: args{[]int{1, 1, 2}}, want: 2},
		{name: "5", args: args{[]int{3, 3, 3}}, want: 3},
		{name: "6", args: args{[]int{1, 2, -2147483648}}, want: -2147483648},
		{name: "7", args: args{[]int{1, -2147483648, 2}}, want: -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
