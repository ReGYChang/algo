package main

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	s := []int{root.Val}
	for i := 0; i < len(root.Children); i++ {
		tmp := preorder(root.Children[i])
		if tmp != nil {
			s = append(s, tmp...)
		}
	}

	return s
}

//Runtime: 3 ms, faster than 75.23% of Go online submissions for N-ary Tree Preorder Traversal.
//Memory Usage: 6.2 MB, less than 34.38% of Go online submissions for N-ary Tree Preorder Traversal.
