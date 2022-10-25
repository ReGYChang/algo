package main

import (
	"testing"

	"algo/structure"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "101-1",
			args: args{[]int{1, 2, 2, 3, 4, 4, 3}},
			want: true,
		},
		{
			name: "101-2",
			args: args{[]int{1, 2, 2, structure.NULL, 3, structure.NULL, 3}},
			want: false,
		},
		{
			name: "101-3",
			args: args{[]int{1, 0}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structure.Int2TreeNode(tt.args.tree)
			if got := isSymmetric(root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
