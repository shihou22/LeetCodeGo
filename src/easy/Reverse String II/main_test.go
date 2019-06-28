package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// 	{name: "1", args: args{"abcdefg", 2}, want: "bacdfeg"},
		// 	{name: "2", args: args{"ba", 1}, want: "ba"},
		// 	{name: "3", args: args{"a", 2}, want: "a"},
		// 	{name: "4", args: args{"abcdefghijklmnopqrstuvwxyz", 2}, want: "bacdfeghjiklnmoprqstvuwxzy"},
		// 	{name: "5", args: args{"abcdefghijklmnopqrstuvwxyz", 3}, want: "cbadefihgjklonmpqrutsvwxzy"},
		// 	{name: "6", args: args{"abcde", 3}, want: "cbade"},
		// 	{name: "7", args: args{"abcdef", 3}, want: "cbadef"},
		{name: "8", args: args{"123456789", 3}, want: "321456987"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
