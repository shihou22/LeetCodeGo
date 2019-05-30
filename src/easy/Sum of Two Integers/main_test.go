package main

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{2, 1}, want: 3},
		{name: "2", args: args{-2, 3}, want: 1},
		{name: "3", args: args{-2, 1}, want: -1},
		{name: "4", args: args{2, 2}, want: 4},
		{name: "5", args: args{2, 3}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
