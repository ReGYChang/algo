package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Sorted Array to Binary Search Tree.
//Memory Usage: 3.4 MB, less than 100.00% of Go online submissions for Convert Sorted Array to Binary Search Tree.

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "108-1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: []int{0, -3, 9, -10, structure.NULL, 5},
		},
		{
			name: "108-2",
			args: args{
				nums: []int{1, 3},
			},
			want: []int{3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantTree := structure.Int2TreeNode(tt.want)
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, wantTree)
			}
		})
	}
}
