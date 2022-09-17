package main

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorderRecursive(root *Node) []int {
	if root == nil {
		return nil
	}

	s := []int{root.Val}
	for i := 0; i < len(root.Children); i++ {
		tmp := preorderRecursive(root.Children[i])
		if tmp != nil {
			s = append(s, tmp...)
		}
	}

	return s
}

//Runtime: 3 ms, faster than 75.23% of Go online submissions for N-ary Tree Preorder Traversal.
//Memory Usage: 6.2 MB, less than 34.38% of Go online submissions for N-ary Tree Preorder Traversal.

func preorderIterative(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{root}
	result := make([]int, 0)

	for i := len(stack); i > 0; {
		result = append(result, stack[i-1].Val)
		for j := len(stack[0].Children); j > 0; j-- {
			stack = append(stack, stack[0].Children[j-1])
		}
		stack = stack[1:]
	}

	return result
}
