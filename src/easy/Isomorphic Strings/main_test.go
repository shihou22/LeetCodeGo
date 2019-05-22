package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"egg", "add"}, want: true},
		{name: "2", args: args{"foo", "bar"}, want: false},
		{name: "3", args: args{"paper", "title"}, want: true},
		{name: "4", args: args{"", ""}, want: true},
		{name: "5", args: args{"a", "bb"}, want: false},
		{name: "5", args: args{"a", "b"}, want: true},
		{name: "6", args: args{"ab", "dc"}, want: true},
		{name: "7", args: args{"ab", "aa"}, want: false},
		{name: "8", args: args{"ab", "aa"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
