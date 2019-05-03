package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{123}, want: 321},
		{name: "2", args: args{-123}, want: -321},
		{name: "3", args: args{120}, want: 21},
		{name: "4", args: args{-4567}, want: -7654},
		{name: "5", args: args{1534236469}, want: 0},
		{name: "6", args: args{7463847412}, want: 2147483647},
		{name: "7", args: args{2147483648}, want: 0},
		{name: "8", args: args{-8463847412}, want: -2147483648},
		{name: "9", args: args{-2147483649}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
