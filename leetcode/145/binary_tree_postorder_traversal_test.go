package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
//Memory Usage: 1.9 MB, less than 90.24% of Go online submissions for Binary Tree Postorder Traversal.

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "145-1",
			args: args{
				tree: []int{1, structure.NULL, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "145-2",
			args: args{
				tree: []int{},
			},
			want: nil,
		},
		{
			name: "145-3",
			args: args{
				tree: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := postorderTraversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_postorderTraversal(b *testing.B) {
	tree := []int{4, 2, 5, 1, 6, 3, 7}
	root := structure.Int2TreeNode(tree)

	for i := 0; i < b.N; i++ {
		postorderTraversal(root)
	}
}
