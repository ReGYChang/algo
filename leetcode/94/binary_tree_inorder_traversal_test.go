package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
//Memory Usage: 2.1 MB, less than 27.93% of Go online submissions for Binary Tree Inorder Traversal.

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "94-1",
			args: args{
				tree: []int{1, structure.NULL, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "94-2",
			args: args{
				tree: []int{},
			},
			want: nil,
		},
		{
			name: "94-3",
			args: args{
				tree: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := inorderTraversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_inorderTraversal(b *testing.B) {
	tree := []int{4, 2, 5, 1, 6, 3, 7}
	root := structure.Int2TreeNode(tree)

	for i := 0; i < b.N; i++ {
		inorderTraversal(root)
	}
}
