package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTStack(root *TreeNode) bool {
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

func isValidBSTMorris(root *TreeNode) bool {
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
