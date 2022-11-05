package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 20 ms, faster than 92.46% of Go online submissions for Two Sum IV - Input is a BST.
//Memory Usage: 6.9 MB, less than 90.16% of Go online submissions for Two Sum IV - Input is a BST.

func Test_searchBST(t *testing.T) {
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
			name: "700-1",
			args: args{
				tree: []int{4, 2, 7, 1, 3},
				val:  2,
			},
			want: []int{2, 1, 3},
		},
		{
			name: "700-2",
			args: args{
				tree: []int{4, 2, 7, 1, 3},
				val:  5,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2TreeNode(tt.args.tree)
			want := structure.Int2TreeNode(tt.want)
			if got := searchBST(head, tt.args.val); !reflect.DeepEqual(got, want) {
				t.Errorf("searchBST() = %v, want %v", got, want)
			}
		})
	}
}
