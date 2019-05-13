package main

import (
	"testing"
)

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3}, want: 0},
		{name: "2", args: args{5}, want: 1},
		{name: "3", args: args{30}, want: 7},
		{name: "4", args: args{1000}, want: 249},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trailingZeroesOther(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3}, want: 0},
		{name: "2", args: args{5}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroesOther(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroesOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trailingZeroesBrute(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{3}, want: 0},
		{name: "2", args: args{5}, want: 1},
		{name: "3", args: args{30}, want: 7},
		{name: "4", args: args{1000}, want: 249},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroesBrute(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroesBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}
