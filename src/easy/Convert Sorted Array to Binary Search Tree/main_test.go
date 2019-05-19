package main

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// {name: "1", args: args{[]int{-10, -3, 0, 5, 9}}, want: &TreeNode{0, &TreeNode{-3, &TreeNode{-10, nil, nil}, nil}, &TreeNode{-9, &TreeNode{5, nil, nil}, nil}}},
		{name: "2", args: args{[]int{3, 5, 8}}, want: &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{8, nil, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
