package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	s := make([]*ListNode, 0)
	for n := head; n != nil; n = n.Next {
		s = append(s, n)
	}

	for i := len(s) - 1; i > 0; i-- {
		s[i].Next = s[i-1]
	}
	s[0].Next = nil

	return s[len(s)-1]
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
//Memory Usage: 2.7 MB, less than 29.35% of Go online submissions for Reverse Linked List.
