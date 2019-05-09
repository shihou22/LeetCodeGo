package main

import (
	"testing"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1}, want: "1"},
		{name: "2", args: args{2}, want: "11"},
		{name: "3", args: args{3}, want: "21"},
		{name: "4", args: args{4}, want: "1211"},
		{name: "5", args: args{5}, want: "111221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countAndSayOther(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1}, want: "1"},
		{name: "2", args: args{2}, want: "11"},
		{name: "3", args: args{3}, want: "21"},
		{name: "4", args: args{4}, want: "1211"},
		{name: "5", args: args{5}, want: "111221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSayOther(tt.args.n); got != tt.want {
				t.Errorf("countAndSayOther() = %v, want %v", got, tt.want)
			}
		})
	}
}
