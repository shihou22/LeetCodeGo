package main

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"ab#c", "ad#c"}, want: true},
		{name: "2", args: args{"ab##", "c#d#"}, want: true},
		{name: "3", args: args{"a##c", "#a#c"}, want: true},
		{name: "4", args: args{"a#c", "b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
