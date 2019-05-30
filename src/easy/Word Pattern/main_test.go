package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"abba", "dog cat cat dog"}, want: true},
		{name: "2", args: args{"abba", "dog cat cat fish"}, want: false},
		{name: "3", args: args{"aaaa", "dog cat cat dog"}, want: false},
		{name: "4", args: args{"abba", "dog dog dog dog"}, want: false},
		{name: "5", args: args{"jquery", "jquery"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
