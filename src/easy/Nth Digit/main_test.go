package main

import (
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 12345678910111213141516171819202122232425262728293031323334353637383940414243444546474849505152535455
		// {name: "1", args: args{3}, want: 3},
		// {name: "2", args: args{11}, want: 0},
		{name: "3", args: args{100}, want: 5},
		{name: "3", args: args{110}, want: 5},
		{name: "4", args: args{1000}, want: 3},
		{name: "5", args: args{10000}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
