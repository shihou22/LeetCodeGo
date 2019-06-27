package main

import (
	"reflect"
	"testing"
)

func Test_findOcurrences(t *testing.T) {
	type args struct {
		text   string
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1", args: args{"alice is a good girl she is a good student", "a", "good"}, want: []string{"girl", "student"}},
		{name: "2", args: args{"we will we will rock you", "we", "will"}, want: []string{"we", "rock"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOcurrences(tt.args.text, tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOcurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
