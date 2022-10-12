package main

import (
	"reflect"
	"testing"
)

//Runtime: 19 ms, faster than 53.81% of Go online submissions for Reshape the Matrix.
//Memory Usage: 6.3 MB, less than 72.03% of Go online submissions for Reshape the Matrix.

func Test_matrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "566-1",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   1,
				c:   4,
			},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "566-2",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   2,
				c:   4,
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "566-3",
			args: args{
				mat: [][]int{{1, 2}, {3, 4}},
				r:   4,
				c:   1,
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
