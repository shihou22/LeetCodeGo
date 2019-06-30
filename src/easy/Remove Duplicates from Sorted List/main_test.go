package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {name: "1", args: args{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}}, want: &ListNode{1, &ListNode{2, nil}}},
		// {name: "2", args: args{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}}, want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
		{name: "3", args: args{&ListNode{1, &ListNode{1, &ListNode{1, nil}}}}, want: &ListNode{1, nil}},
		{name: "4", args: args{&ListNode{1, &ListNode{1, nil}}}, want: &ListNode{1, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
