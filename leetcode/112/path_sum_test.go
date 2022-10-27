package main

import (
	"testing"

	"algo/structure"
)

//Runtime: 4 ms, faster than 92.41% of Go online submissions for Path Sum.
//Memory Usage: 4.7 MB, less than 78.45% of Go online submissions for Path Sum.

func Test_hasPathSum(t *testing.T) {
	type args struct {
		tree      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "112-1",
			args: args{
				tree:      []int{5, 4, 8, 11, structure.NULL, 13, 4, 7, 2, structure.NULL, structure.NULL, structure.NULL, 1},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "112-2",
			args: args{
				tree:      []int{1, 2, 3},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "112-3",
			args: args{
				tree:      []int{},
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "112-4",
			args: args{
				tree:      []int{1, 2},
				targetSum: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := hasPathSum(root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
