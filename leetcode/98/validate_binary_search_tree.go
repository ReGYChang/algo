package main

import "algo/structure"

func isValidBSTStack(root *structure.TreeNode) bool {
	stack := make([]*structure.TreeNode, 0)
	cur := root
	previous := -2 << 31
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val <= previous {
			return false
		}
		previous = node.Val
		cur = node.Right
	}
	return true
}

func isValidBSTMorris(root *structure.TreeNode) bool {
	cur, pre := root, root
	preVal := -2 << 31
	for cur != nil {
		if cur.Left == nil {
			if cur.Val <= preVal {
				return false
			}
			preVal = cur.Val

			cur = cur.Right
		} else {
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				if cur.Val <= preVal {
					return false
				}
				preVal = cur.Val

				pre.Right = nil
				cur = cur.Right
			}
		}
	}
	return true
}
