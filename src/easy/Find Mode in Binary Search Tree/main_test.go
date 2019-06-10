package main

import (
	"reflect"
	"testing"
)

func Test_findModeMap(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{&TreeNode{1, nil, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}}, want: []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findModeMap(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findModeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
