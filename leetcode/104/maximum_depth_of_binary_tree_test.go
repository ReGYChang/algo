package main

import (
	"testing"

	"algo/structure"
)

//Runtime: 8 ms, faster than 82.43% of Go online submissions for Maximum Depth of Binary Tree.
//Memory Usage: 4.3 MB, less than 62.31% of Go online submissions for Maximum Depth of Binary Tree.

func Test_maxDepth(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "104-1",
			args: args{tree: []int{3, 9, 20, structure.NULL, structure.NULL, 15, 7}},
			want: 3,
		},
		{
			name: "104-2",
			args: args{tree: []int{1, structure.NULL, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
