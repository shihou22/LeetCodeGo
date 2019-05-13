package main

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1}, want: "A"},
		{name: "2", args: args{26}, want: "Z"},
		{name: "3", args: args{27}, want: "AA"},
		{name: "4", args: args{28}, want: "AB"},
		{name: "5", args: args{701}, want: "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
