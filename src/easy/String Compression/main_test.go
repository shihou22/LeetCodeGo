package main

import (
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}, want: 6},
		// {name: "2", args: args{[]byte{'a'}}, want: 1},
		{name: "3", args: args{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hoge(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hoge(tt.args.a)
		})
	}
}
