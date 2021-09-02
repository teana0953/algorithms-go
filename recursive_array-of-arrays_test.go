package main

import (
	"reflect"
	"testing"
)

func TestArrayOfArrays(t *testing.T) {
	type args struct {
		ary []interface{}
		acc []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ary: []interface{}{
					[]interface{}{
						"a",
						"b",
						[]interface{}{
							"c",
							[]interface{}{
								[]interface{}{
									"d",
								},
							},
						},
						"e",
						"f",
					},
				},
			},
			want: []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayOfArrays(tt.args.ary, tt.args.acc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
