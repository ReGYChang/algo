package main

import "algo/structure"

func detectCycle(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*structure.ListNode]int16)
	for head.Next != nil {
		if _, exist := m[head]; exist {
			return head
		}
		m[head] = 1
		head = head.Next
	}

	return nil
}
