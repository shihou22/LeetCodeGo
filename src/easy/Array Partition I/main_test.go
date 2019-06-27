package main

import "testing"

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{[]int{1, 4, 3, 2}}, want: 4},
		{name: "2", args: args{[]int{7, 3, 1, 0, 0, 6}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
