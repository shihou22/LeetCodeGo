package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{&ListNode{2, &ListNode{0, &ListNode{4, &ListNode{-1, nil}}}}}, want: true},
		{name: "1", args: args{&ListNode{1, &ListNode{2, nil}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
