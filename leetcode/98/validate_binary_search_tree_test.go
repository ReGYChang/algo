package main

import (
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Validate Binary Search Tree.
//Memory Usage: 5.2 MB, less than 85.09% of Go online submissions for Validate Binary Search Tree.

func Test_isValidBSTStack_98(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validate binary search tree",
			args: args{
				input: []int{10, 5, 15, structure.NULL, structure.NULL, 6, 20},
			},
			want: false,
		},
		{
			name: "validate binary search tree 2",
			args: args{
				input: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "validate binary search tree 3",
			args: args{
				input: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "validate binary search tree 4",
			args: args{
				input: []int{5, 1, 4, structure.NULL, structure.NULL, 3, 6},
			},
			want: false,
		},
		{
			name: "validate binary search tree 5",
			args: args{
				input: []int{2, 2, 2, 2, 2, 2, 2},
			},
			want: false,
		},
		{
			name: "validate binary search tree 6",
			args: args{
				input: []int{2, 2, 2},
			},
			want: false,
		},
		{
			name: "validate binary search tree 7",
			args: args{
				input: []int{2},
			},
			want: true,
		},
		{
			name: "validate binary search tree 8",
			args: args{
				input: []int{3, structure.NULL, 30, 10, structure.NULL, structure.NULL, 15, structure.NULL, 45},
			},
			want: false,
		},
		{
			name: "validate binary search tree 9",
			args: args{
				input: []int{0, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.input)
			if got := isValidBSTStack(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 7 ms, faster than 73.40% of Go online submissions for Validate Binary Search Tree.
//Memory Usage: 5 MB, less than 96.39% of Go online submissions for Validate Binary Search Tree.

func Test_isValidBSTMorris_98(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validate binary search tree",
			args: args{
				input: []int{10, 5, 15, structure.NULL, structure.NULL, 6, 20},
			},
			want: false,
		},
		{
			name: "validate binary search tree 2",
			args: args{
				input: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "validate binary search tree 3",
			args: args{
				input: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "validate binary search tree 4",
			args: args{
				input: []int{5, 1, 4, structure.NULL, structure.NULL, 3, 6},
			},
			want: false,
		},
		{
			name: "validate binary search tree 5",
			args: args{
				input: []int{2, 2, 2, 2, 2, 2, 2},
			},
			want: false,
		},
		{
			name: "validate binary search tree 6",
			args: args{
				input: []int{2, 2, 2},
			},
			want: false,
		},
		{
			name: "validate binary search tree 7",
			args: args{
				input: []int{2},
			},
			want: true,
		},
		{
			name: "validate binary search tree 8",
			args: args{
				input: []int{3, structure.NULL, 30, 10, structure.NULL, structure.NULL, 15, structure.NULL, 45},
			},
			want: false,
		},
		{
			name: "validate binary search tree 9",
			args: args{
				input: []int{0, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.input)
			if got := isValidBSTMorris(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
