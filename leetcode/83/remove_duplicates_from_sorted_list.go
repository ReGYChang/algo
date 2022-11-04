package main

import "algo/structure"

type ListNode = structure.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = deDuplicates(head.Next, head.Val)
	return head
}

func deDuplicates(root *ListNode, preVal int) *ListNode {
	if root == nil {
		return nil
	}
	if preVal == root.Val {
		return deDuplicates(root.Next, root.Val)
	}
	root.Next = deDuplicates(root.Next, root.Val)
	return root
}