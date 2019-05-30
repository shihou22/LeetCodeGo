package main

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "1", args: args{"abcd", "abcde"}, want: 'e'},
		{name: "2", args: args{"abcd", "eabcd"}, want: 'e'},
		{name: "3", args: args{"abcd", "eabecde"}, want: 'e'},
		{name: "4", args: args{"a", "aa"}, want: 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
