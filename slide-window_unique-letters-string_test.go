package main

import "testing"

func TestUniqueLetterString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				str: "thisisshowwedoit",
			},
			want: 6,
		},
		{
			name: "test1",
			args: args{
				str: "GEEKSFORGEEKS",
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				str: "aaaaaaaa",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueLetterString(tt.args.str); got != tt.want {
				t.Errorf("UniqueLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
