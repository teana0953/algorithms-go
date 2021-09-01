package main

import (
	"reflect"
	"testing"
)

func TestAveragePair(t *testing.T) {
	type args struct {
		sortedAry []float32
		target    float32
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]float32
	}{
		// TODO: Add test cases.
		{
			name: "have AveragePair",
			args: args{
				sortedAry: []float32{-11, 0, 1, 2, 3, 9, 14, 17, 21},
				target:    1.5,
			},
			wantResult: [][]float32{{-11, 14}, {0, 3}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AveragePair(tt.args.sortedAry, tt.args.target); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AveragePair() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
