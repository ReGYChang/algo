package main

import (
	"reflect"
	"testing"
)

var NULL = -1 << 63

//Runtime: 41 ms, faster than 58.20% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
//Memory Usage: 7.2 MB, less than 80.32% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root []int
		p    []int
		q    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lowest common ancestor",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
				p:    []int{2},
				q:    []int{8},
			},
			want: 6,
		},
		{
			name: "lowest common ancestor 2",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
				p:    []int{2},
				q:    []int{4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := Ints2TreeNode(tt.args.root)
			p := Ints2TreeNode(tt.args.p)
			q := Ints2TreeNode(tt.args.q)
			if got := lowestCommonAncestor(root, p, q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
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
