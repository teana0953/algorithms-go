package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		newStr string
		oldStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "empty ok",
			args: args{
				newStr: "",
				oldStr: "hello Dear",
			},
			want: true,
		},
		{
			name: "hello Dear",
			args: args{
				newStr: "hello",
				oldStr: "hello Dear",
			},
			want: true,
		},
		{
			name: "brooklyn",
			args: args{
				newStr: "book",
				oldStr: "brooklyn",
			},
			want: true,
		},
		{
			name: "order matter",
			args: args{
				newStr: "abc",
				oldStr: "bac",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.newStr, tt.args.oldStr); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
