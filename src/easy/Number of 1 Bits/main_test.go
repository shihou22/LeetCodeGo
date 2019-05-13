package main

import (
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{00000000000000000000000000001011}, want: 3},
		{name: "2", args: args{00000000000000000000000010000000}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeightOther(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{00000000000000000000000000001011}, want: 3},
		{name: "2", args: args{00000000000000000000000010000000}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeightOther(tt.args.num); got != tt.want {
				t.Errorf("hammingWeightOther() = %v, want %v", got, tt.want)
			}
		})
	}
}
