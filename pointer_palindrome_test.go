package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "false",
			args: args{
				str: "awesome",
			},
			want: false,
		},
		{
			name: "false",
			args: args{
				str: "foobar",
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				str: "tacocat",
			},
			want: true,
		},
		{
			name: "true",
			args: args{
				str: "amanaplanacanalpanama",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.str); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
