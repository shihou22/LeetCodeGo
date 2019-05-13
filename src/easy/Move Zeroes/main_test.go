package main

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{0, 1, 0, 3, 12}}}, //1,3,12,0,0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}

func Test_moveZeroesOther(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{0, 1, 0, 3, 12}}}, //1,3,12,0,0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesOther(tt.args.nums)
		})
	}
}

func Test_moveZeroesBruteForce(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{0, 1, 0, 3, 12}}}, //1,3,12,0,0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesBruteForce(tt.args.nums)
		})
	}
}
