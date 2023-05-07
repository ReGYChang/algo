package main

import "algo/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *structure.ListNode, n int) *structure.ListNode {
	s := make([]*structure.ListNode, 0)
	for node := head; node != nil; node = node.Next {
		s = append(s, node)
	}
	s[len(s)-n] = nil

	res := &structure.ListNode{}
	t := res
	for i := 0; i < len(s); i++ {
		if s[i] != nil {
			t.Next = &structure.ListNode{Val: s[i].Val}
			t = t.Next
		} else {
			if i < len(s)-1 {
				t.Next = &structure.ListNode{Val: s[i+1].Val}
				t = t.Next
				i++
			}
		}
	}

	return res.Next
}
