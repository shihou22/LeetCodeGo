package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{[]string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "2", args: args{[]string{"dog", "racecar", "car"}}, want: ""},
		{name: "2", args: args{[]string{}}, want: ""},
		{name: "2", args: args{[]string{""}}, want: ""},
		{name: "2", args: args{[]string{"aac", "a", "ccc"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
