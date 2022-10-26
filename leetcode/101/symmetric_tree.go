package main

import "algo/structure"

type TreeNode = structure.TreeNode

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
}
