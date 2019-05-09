package main

import (
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"abab"}, want: true},
		{name: "2", args: args{"aba"}, want: false},
		{name: "3", args: args{"abcabcabcabc"}, want: true},
		{name: "4", args: args{"ababab"}, want: true},
		{name: "5", args: args{"aaa"}, want: true},
		{name: "6", args: args{"a"}, want: false},
		{name: "7", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatedSubstringPatternOther(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {name: "1", args: args{"abab"}, want: true},
		// {name: "2", args: args{"aba"}, want: false},
		// {name: "3", args: args{"abcabcabcabc"}, want: true},
		// {name: "4", args: args{"ababab"}, want: true},
		{name: "5", args: args{"aaa"}, want: true},
		// {name: "6", args: args{"a"}, want: false},
		// {name: "7", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPatternOther(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPatternOther() = %v, want %v", got, tt.want)
			}
		})
	}
}
