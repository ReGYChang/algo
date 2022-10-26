package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
//Memory Usage: 2.2 MB, less than 70.21% of Go online submissions for Invert Binary Tree.

func Test_invertTree(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "226-1",
			args: args{tree: []int{4, 2, 7, 1, 3, 6, 9}},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "226-2",
			args: args{tree: []int{2, 1, 3}},
			want: []int{2, 3, 1},
		},
		{
			name: "226-3",
			args: args{tree: []int{}},
			want: []int{},
		},
		{
			name: "226-4",
			args: args{tree: []int{1, 2, structure.NULL}},
			want: []int{1, structure.NULL, 2},
		},
		{
			name: "226-5",
			args: args{tree: []int{2, 3, structure.NULL, 1}},
			want: []int{2, structure.NULL, 3, structure.NULL, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			want := structure.Int2TreeNode(tt.want)
			if got := invertTree(root); !reflect.DeepEqual(got, want) {
				t.Errorf("invertTree() = %v, want %v", got, want)
			}
		})
	}
}
