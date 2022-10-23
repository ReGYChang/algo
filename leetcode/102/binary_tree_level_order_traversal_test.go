package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 2 ms, faster than 67.53% of Go online submissions for Binary Tree Level Order Traversal.
//Memory Usage: 2.7 MB, less than 89.92% of Go online submissions for Binary Tree Level Order Traversal.

func Test_levelOrder_102(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "binary_tree_level_order_traversal",
			args: args{
				input: []int{},
			},
			want: [][]int{},
		},
		{
			name: "binary_tree_level_order_traversal_2",
			args: args{
				input: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name: "binary_tree_level_order_traversal_3",
			args: args{
				input: []int{3, 9, 20, structure.NULL, structure.NULL, 15, 7},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.input)
			if got := levelOrder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
