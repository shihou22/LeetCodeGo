package main

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{27}, want: true},
		{name: "2", args: args{0}, want: false},
		{name: "3", args: args{1}, want: true},
		{name: "4", args: args{29}, want: false},
		{name: "5", args: args{9}, want: true},
		{name: "6", args: args{45}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
