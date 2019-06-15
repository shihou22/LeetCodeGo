package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"leetcode"}, want: 0},
		{name: "2", args: args{"loveleetcode"}, want: 2},
		{name: "3", args: args{""}, want: -1},
		{name: "4", args: args{"cc"}, want: -1},
		{name: "5", args: args{"aadadaad"}, want: -1},
		{name: "6", args: args{"z"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
