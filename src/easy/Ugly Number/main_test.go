package main

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{6}, want: true},
		{name: "2", args: args{8}, want: true},
		{name: "3", args: args{14}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
