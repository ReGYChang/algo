package main

import "algo/structure"

func removeElements(head *structure.ListNode, val int) *structure.ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}

	head.Next = removeElements(head.Next, val)
	return head
}
