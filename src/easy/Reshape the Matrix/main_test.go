package main

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		nums [][]int
		r    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "1", args: args{[][]int{{1, 2}, {3, 4}}, 1, 4}, want: [][]int{{1, 2, 3, 4}}},
		{name: "2", args: args{[][]int{{1, 2}, {3, 4}}, 2, 4}, want: [][]int{{1, 2}, {3, 4}}},
		{name: "3", args: args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}, 42, 5}, want: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.nums, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
