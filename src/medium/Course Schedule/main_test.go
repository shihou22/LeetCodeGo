package main

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{2, [][]int{{1, 0}}}, want: true},
		{name: "2", args: args{2, [][]int{{1, 0}, {0, 1}}}, want: false},
		{name: "3", args: args{3, [][]int{{0, 2}, {1, 2}, {2, 0}}}, want: false},
		{name: "4", args: args{3, [][]int{{1, 0}, {2, 1}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
