package main

import (
	"reflect"
	"testing"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
//Memory Usage: 1.9 MB, less than 86.92% of Go online submissions for Pascal's Triangle.

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "118-1",
			args: args{5},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "118-2",
			args: args{1},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
