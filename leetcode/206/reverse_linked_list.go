package main

import "algo/structure"

func reverseList(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return nil
	}

	s := make([]*structure.ListNode, 0)
	for n := head; n != nil; n = n.Next {
		s = append(s, n)
	}

	for i := len(s) - 1; i > 0; i-- {
		s[i].Next = s[i-1]
	}
	s[0].Next = nil

	return s[len(s)-1]
}
