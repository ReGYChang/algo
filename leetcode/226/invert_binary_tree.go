package main

import "algo/structure"

type TreeNode = structure.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invert(root.Left, root.Right)
	return root
}

func invert(left, right *TreeNode) (*TreeNode, *TreeNode) {
	if left == nil && right == nil {
		return nil, nil
	}
	if left != nil {
		left.Left, left.Right = invert(left.Left, left.Right)
	}
	if right != nil {
		right.Left, right.Right = invert(right.Left, right.Right)
	}

	return right, left
}
