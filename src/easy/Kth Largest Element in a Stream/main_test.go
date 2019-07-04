package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestKthLargest_Add(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *KthLargest
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Add(tt.args.val); got != tt.want {
				t.Errorf("KthLargest.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
