package main

import "algo/structure"

func preorderTraversal(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}
