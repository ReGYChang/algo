package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Search a 2D Matrix.
//Memory Usage: 2.7 MB, less than 78.10% of Go online submissions for Search a 2D Matrix.

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "74-1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{12, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "74-2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
