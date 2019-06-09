package main

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{"5F3Z-2e-9-w", 4}, want: "5F3Z-2E9W"},
		{name: "2", args: args{"2-5g-3-J", 2}, want: "2-5G-3J"},
		{name: "3", args: args{"2-4A0r7-4k", 4}, want: "24A0-R74K"},
		{name: "4", args: args{"2-4A0r7-4k", 3}, want: "24-A0R-74K"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
