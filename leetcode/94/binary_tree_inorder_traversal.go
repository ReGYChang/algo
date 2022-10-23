package main

import "algo/structure"

func inorderTraversal(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}
