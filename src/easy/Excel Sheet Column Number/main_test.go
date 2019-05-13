package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"A"}, want: 1},
		{name: "2", args: args{"Z"}, want: 26},
		{name: "3", args: args{"AA"}, want: 27},
		{name: "4", args: args{"AB"}, want: 28},
		{name: "5", args: args{"ZY"}, want: 701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
