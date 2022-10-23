package main

import "algo/structure"

func middleNode(head *structure.ListNode) *structure.ListNode {
	p := head
	count := 1

	for p.Next != nil {
		count++
		p = p.Next
	}

	count = count / 2
	for i := 0; i < count; i++ {
		head = head.Next
	}

	return head
}
