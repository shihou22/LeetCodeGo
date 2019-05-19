package main

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}}, want: true},
		{name: "2", args: args{&TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, nil}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
