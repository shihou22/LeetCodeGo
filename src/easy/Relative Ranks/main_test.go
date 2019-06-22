package main

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{[]int{5, 1, 3, 2, 4}}, want: []string{"Gold Medal", "5", "Bronze Medal", "4", "Silver Medal"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
