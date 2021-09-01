package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	type args struct {
		ary1 []int
		ary2 []int
	}
	tests := []struct {
		name        string
		args        args
		wantResults []int
	}{
		// TODO: Add test cases.
		{
			name: "has intersection",
			args: args{
				ary1: []int{1, 3, 6, 7, 9, 100},
				ary2: []int{2, 5, 6, 100, 9},
			},
			wantResults: []int{6, 9, 100},
		},
		{
			name: "no intersection",
			args: args{
				ary1: []int{1, 3, 8, 7, 20, 99},
				ary2: []int{2, 5, 6, 100, 9},
			},
			wantResults: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := FindIntersection(tt.args.ary1, tt.args.ary2); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("FindIntersection() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
