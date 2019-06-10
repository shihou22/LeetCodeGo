package main

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {name: "1", args: args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, want: []int{-1, 3, -1}},
		// {name: "2", args: args{[]int{2, 4}, []int{1, 2, 3, 4}}, want: []int{3, -1}},
		{name: "3", args: args{[]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}}, want: []int{7, 7, 7, 7, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
