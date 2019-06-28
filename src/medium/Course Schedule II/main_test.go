package main

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{2, [][]int{{1, 0}}}, want: []int{0, 1}},
		{name: "2", args: args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, want: []int{0, 1, 2, 3}},
		{name: "3", args: args{3, [][]int{{1, 0}, {1, 2}, {0, 1}}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
