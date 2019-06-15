package main

import (
	"reflect"
	"testing"
)

func Test_sortAWithB(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{name: "1", want: []int{-23, 24, 56, 74, 87, 91}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAWithB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAWithB() = %v, want %v", got, tt.want)
			}
		})
	}
}
