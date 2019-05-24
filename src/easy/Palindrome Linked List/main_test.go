package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{&ListNode{1, &ListNode{2, nil}}}, want: false},
		{name: "2", args: args{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}}, want: true},
		{name: "3", args: args{nil}, want: true},
		{name: "4", args: args{&ListNode{1, &ListNode{3, &ListNode{0, &ListNode{2, nil}}}}}, want: false},
		{name: "5", args: args{&ListNode{1, &ListNode{2, &ListNode{0, &ListNode{2, nil}}}}}, want: false},
		{name: "6", args: args{&ListNode{1, &ListNode{0, &ListNode{1, nil}}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkSlice()
		})
	}
}
