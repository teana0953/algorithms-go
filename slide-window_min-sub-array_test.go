package main

import "testing"

func TestMinSubLength(t *testing.T) {
	type args struct {
		ary []int
		sum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "4",
			args: args{
				ary: []int{8, 1, 6, 15, 3, 16, 5, 7, 14, 30, 12},
				sum: 60,
			},
			want: 4,
		},
		{
			name: "test",
			args: args{
				ary: []int{2, 3, 1, 2, 4, 3},
				sum: 7,
			},
			want: 2, // [4, 3]
		},
		{
			name: "test2",
			args: args{
				ary: []int{1, 4, 4},
				sum: 4,
			},
			want: 1, // [4]
		},
		{
			name: "test3",
			args: args{
				ary: []int{1, 1, 1, 1, 1, 1, 1, 1},
				sum: 11,
			},
			want: 0, // [4]
		},
		{
			name: "test4",
			args: args{
				ary: []int{3, 3, 2, 1},
				sum: 6,
			},
			want: 2, // [3, 3]
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubLength(tt.args.ary, tt.args.sum); got != tt.want {
				t.Errorf("MinSubLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
