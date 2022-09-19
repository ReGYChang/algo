package main

import (
	"reflect"
	"testing"
)

var NULL = -1 << 63

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
				input: []int{3, 9, 20, NULL, NULL, 15, 7},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := Ints2TreeNode(tt.args.input)
			if got := levelOrder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
