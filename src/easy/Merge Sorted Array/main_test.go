package main

import (
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_mergeUsingSort(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeUsingSort(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}
