package main

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"cdadabcc"}, want: "adbc"},
		{name: "2", args: args{"abcd"}, want: "abcd"},
		{name: "3", args: args{"ecbacba"}, want: "eacb"},
		{name: "4", args: args{"leetcode"}, want: "letcod"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.text); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
