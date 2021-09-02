package main

import "testing"

func TestRecursive(t *testing.T) {
	type args struct {
		n int
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
				n: 3,
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recursive(tt.args.n); got != tt.want {
				t.Errorf("Recursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
