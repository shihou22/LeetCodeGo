package main

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{&TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}}, want: []string{"1->2->5", "1->3"}},
		{name: "2", args: args{&TreeNode{1, nil, nil}}, want: []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
