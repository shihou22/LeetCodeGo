package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"hello", "ll"}, want: 2},
		{name: "2", args: args{"aaaaa", "bba"}, want: -1},
		{name: "3", args: args{"", "bba"}, want: -1},
		{name: "4", args: args{"aaa", ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
