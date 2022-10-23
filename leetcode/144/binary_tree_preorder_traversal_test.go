package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 4 ms, faster than 12.62% of Go online submissions for Binary Tree Preorder Traversal.
//Memory Usage: 2.1 MB, less than 40.97% of Go online submissions for Binary Tree Preorder

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "144-1",
			args: args{
				tree: []int{1, structure.NULL, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "144-2",
			args: args{
				tree: []int{},
			},
			want: nil,
		},
		{
			name: "144-3",
			args: args{
				tree: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := preorderTraversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_preorderTraversal(b *testing.B) {
	tree := []int{4, 2, 5, 1, 6, 3, 7}
	root := structure.Int2TreeNode(tree)

	for i := 0; i < b.N; i++ {
		preorderTraversal(root)
	}
}
