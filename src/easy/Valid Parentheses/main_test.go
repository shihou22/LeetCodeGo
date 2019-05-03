package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"()"}, want: true},
		{name: "2", args: args{"()[]{}"}, want: true},
		{name: "3", args: args{"(]"}, want: false},
		{name: "4", args: args{"([)]"}, want: false},
		{name: "5", args: args{"{[]}"}, want: true},
		{name: "6", args: args{"]"}, want: false},
		{name: "7", args: args{""}, want: true},
		{name: "8", args: args{"){"}, want: false},
		{name: "9", args: args{"(("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
