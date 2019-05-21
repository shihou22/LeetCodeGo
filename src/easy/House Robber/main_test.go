package main

import (
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 2, 3, 1}}, want: 4},
		{name: "2", args: args{[]int{2, 7, 9, 3, 1}}, want: 12},
		{name: "3", args: args{[]int{}}, want: 0},
		{name: "4", args: args{[]int{1, 2}}, want: 2},
		{name: "5", args: args{[]int{0}}, want: 0},
		{name: "6", args: args{[]int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robTLE(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 2, 3, 1}}, want: 4},
		{name: "2", args: args{[]int{2, 7, 9, 3, 1}}, want: 12},
		{name: "3", args: args{[]int{}}, want: 0},
		{name: "4", args: args{[]int{1, 2}}, want: 2},
		{name: "5", args: args{[]int{0}}, want: 0},
		{name: "6", args: args{[]int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robTLE(tt.args.nums); got != tt.want {
				t.Errorf("robTLE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robDP1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 2, 3, 1}}, want: 4},
		{name: "2", args: args{[]int{2, 7, 9, 3, 1}}, want: 12},
		{name: "3", args: args{[]int{}}, want: 0},
		{name: "4", args: args{[]int{1, 2}}, want: 2},
		{name: "5", args: args{[]int{0}}, want: 0},
		{name: "6", args: args{[]int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robDP1(tt.args.nums); got != tt.want {
				t.Errorf("robDP1() = %v, want %v", got, tt.want)
			}
		})
	}
}
