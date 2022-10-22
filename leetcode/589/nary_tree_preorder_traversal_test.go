package main

import (
	"reflect"
	"testing"
)

var NULL = -1 << 63

func Test_preorderRecursive_589(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nary tree preorder traversal",
			args: args{
				input: []int{1, NULL, 3, 2, 4, NULL, 5, 6},
			},
			want: []int{1, 3, 5, 6, 2, 4},
		},
		{
			name: "nary tree preorder traversal2",
			args: args{
				input: []int{1, NULL, 2, 3, 4, 5, NULL, NULL, 6, 7, NULL, 8, NULL, 9, 10, NULL, NULL, 11, NULL, 12, NULL, 13, NULL, NULL, 14},
			},
			want: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootOne := int2NaryNode(tt.args.input)
			if got := preorderRecursive(rootOne); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderIterative_589(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nary tree preorder traversal",
			args: args{
				input: []int{1, NULL, 3, 2, 4, NULL, 5, 6},
			},
			want: []int{1, 3, 5, 6, 2, 4},
		},
		{
			name: "nary tree preorder traversal2",
			args: args{
				input: []int{1, NULL, 2, 3, 4, 5, NULL, NULL, 6, 7, NULL, 8, NULL, 9, 10, NULL, NULL, 11, NULL, 12, NULL, 13, NULL, NULL, 14},
			},
			want: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootOne := int2NaryNode(tt.args.input)
			if got := preorderIterative(rootOne); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func int2NaryNode(nodes []int) *Node {
	root := &Node{}
	if len(nodes) > 1 {
		root.Val = nodes[0]
	}
	queue := []*Node{}
	queue = append(queue, root)
	i := 1
	count := 0
	for i < len(nodes) {
		node := queue[0]

		childrens := []*Node{}
		for ; i < len(nodes) && nodes[i] != NULL; i++ {
			tmp := &Node{Val: nodes[i]}
			childrens = append(childrens, tmp)
			queue = append(queue, tmp)
		}
		count++
		if count%2 == 0 {
			queue = queue[1:]
			count = 1
		}
		if node != nil {
			node.Children = childrens
		}
		i++
	}
	return root
}
