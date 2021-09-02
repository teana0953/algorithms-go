package main

import (
	"reflect"
	"testing"
)

func TestMaxSum(t *testing.T) {
	type args struct {
		ary  []int
		size int
	}

	want0 := 12

	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 3,
			},
			want: &want0,
		},
		{
			name: "size > array length",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSum(tt.args.ary, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSumImproved(t *testing.T) {
	type args struct {
		ary  []int
		size int
	}

	want0 := 12

	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 3,
			},
			want: &want0,
		},
		{
			name: "size > array length",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumImproved(tt.args.ary, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSumImproved() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinSum(t *testing.T) {
	type args struct {
		ary  []int
		size int
	}

	want0 := -28

	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 3,
			},
			want: &want0,
		},
		{
			name: "size > array length",
			args: args{
				ary:  []int{2, 7, 3, 0, 6, 1, -5, -12, -11},
				size: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSum(tt.args.ary, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
