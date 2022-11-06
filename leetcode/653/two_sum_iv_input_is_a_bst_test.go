
package main

import (
	"testing"

	"algo/structure"
)

//Runtime: 27 ms, faster than 84.35% of Go online submissions for Two Sum IV - Input is a BST.
//Memory Usage: 7.2 MB, less than 87.15% of Go online submissions for Two Sum IV - Input is a BST.

func Test_findTarget(t *testing.T) {
	type args struct {
		tree []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "653-1",
			args: args{
				tree: []int{5, 3, 6, 2, 4, structure.NULL, 7},
				k:    9,
			},
			want: true,
		},
		{
			name: "653-2",
			args: args{
				tree: []int{5, 3, 6, 2, 4, structure.NULL, 7},
				k:    28,
			},
			want: false,
		},
		{
			name: "653-3",
			args: args{
				tree: []int{1},
				k:    2,
			},
			want: false,
		},
		{
			name: "653-4",
			args: args{
				tree: []int{3, 9, 20, structure.NULL, structure.NULL, 15, 7},
				k:    29,
			},
			want: true,
		},
		{
			name: "653-5",
			args: args{
				tree: []int{1, 2, 3, 4, structure.NULL, structure.NULL, 5},
				k:    9,
			},
			want: true,
		},
		{
			name: "653-6",
			args: args{
				tree: []int{2, 1, 3},
				k:    4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := findTarget(root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
