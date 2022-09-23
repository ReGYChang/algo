package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pNode := root
	previous := -2 << 31
	for len(stack) > 0 || pNode != nil {
		if pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Val <= previous {
				return false
			}
			previous = node.Val
			pNode = node.Right
		}
	}
	return true
}
