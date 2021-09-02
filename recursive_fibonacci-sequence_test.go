package main

import "testing"

func TestFibonacciSequence(t *testing.T) {
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
			want: 2,
		},
		{
			name: "test1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciSequence(tt.args.n); got != tt.want {
				t.Errorf("FibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
