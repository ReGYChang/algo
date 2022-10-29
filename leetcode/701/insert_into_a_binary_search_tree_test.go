package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 25 ms, faster than 95.85% of Go online submissions for Insert into a Binary Search Tree.
//Memory Usage: 7.3 MB, less than 80.83% of Go online submissions for Insert into a Binary Search Tree.

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		tree []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "701-1",
			args: args{
				tree: []int{4, 2, 7, 1, 3},
				val:  5,
			},
			want: []int{4, 2, 7, 1, 3, 5},
		},
		{
			name: "701-2",
			args: args{
				tree: []int{40, 20, 60, 10, 30, 50, 70},
				val:  25,
			},
			want: []int{40, 20, 60, 10, 30, 50, 70, structure.NULL, structure.NULL, 25},
		},
		{
			name: "701-3",
			args: args{
				tree: []int{4, 2, 7, 1, 3, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL, structure.NULL},
				val:  5,
			},
			want: []int{4, 2, 7, 1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			want := structure.Int2TreeNode(tt.want)
			if got := insertIntoBST(root, tt.args.val); !reflect.DeepEqual(got, want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, want)
			}
		})
	}
}
