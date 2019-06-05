package main

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{2, 7, 4, 1, 8, 1}}, want: 1},
		{name: "2", args: args{[]int{5, 5, 4, 4, 3, 2}}, want: 1},
		{name: "3", args: args{[]int{7, 6, 7, 6, 9}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
