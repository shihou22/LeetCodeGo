package main

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{&TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{-2, nil, nil}}, &TreeNode{2, nil, &TreeNode{1, nil, nil}}}, &TreeNode{-3, nil, &TreeNode{11, nil, nil}}}, 8}, want: 3},
		{name: "2", args: args{&TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{4, nil, nil}}}, 22}, want: 3},
		{name: "3", args: args{&TreeNode{1, nil, nil}, 0}, want: 0},
		{name: "4", args: args{&TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}}}, 3}, want: 2},
		{name: "5", args: args{&TreeNode{0, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}, 1}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
