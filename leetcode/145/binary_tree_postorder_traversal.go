package main

import "algo/structure"

func postorderTraversal(root *structure.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := postorderTraversal(root.Left)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)

	return res
}
