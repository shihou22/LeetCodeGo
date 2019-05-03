package main

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{16}, want: true},
		{name: "2", args: args{5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
