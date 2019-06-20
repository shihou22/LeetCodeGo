package main

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{[]int{1, 2, 3}, []int{2}}, want: 1},
		// {name: "2", args: args{[]int{1, 2, 3, 4}, []int{1, 4}}, want: 1},
		// {name: "3", args: args{[]int{1, 5}, []int{2}}, want: 3},
		// {name: "4", args: args{[]int{1, 5}, []int{6}}, want: 5},
		// {name: "5", args: args{[]int{1, 5}, []int{1}}, want: 4},
		// {name: "6", args: args{[]int{1, 5}, []int{0, 2, 4}}, want: 1},
		// {name: "7", args: args{[]int{1, 2, 3, 5, 15}, []int{2, 30}}, want: 13},
		{name: "8", args: args{[]int{1, 2, 3, 5, 15}, []int{1, 15}}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
