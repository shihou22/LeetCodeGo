package main

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{2.00000, 10}, want: 1024.00000},
		{name: "2", args: args{2.10000, 3}, want: 9.26100},
		{name: "3", args: args{2.00000, -2}, want: 0.25000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
