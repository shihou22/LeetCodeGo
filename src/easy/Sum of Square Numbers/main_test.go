package main

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{5}, want: true},
		{name: "2", args: args{3}, want: false},
		{name: "3", args: args{4}, want: true},
		{name: "4", args: args{0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
