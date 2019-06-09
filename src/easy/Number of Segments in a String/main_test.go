package main

import "testing"

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"Hello, my name is John"}, want: 5},
		{name: "2", args: args{""}, want: 0},
		{name: "3", args: args{"Hello"}, want: 1},
		{name: "4", args: args{"                "}, want: 0},
		{name: "5", args: args{"Of all the gin joints in all the towns in all the world,   "}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
