package main

import "testing"

func Test_SameFrequency(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "same length",
			args: args{
				str1: "abbc",
				str2: "aabc",
			},
			want: false,
		},
		{
			name: "same length",
			args: args{
				str1: "abba",
				str2: "abab",
			},
			want: true,
		},
		{
			name: "diff length",
			args: args{
				str1: "aasdebasdf",
				str2: "adfeebed",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameFrequency(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("sameFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
