package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]int16)
	for head.Next != nil {
		if _, exist := m[head]; exist {
			return head
		}
		m[head] = 1
		head = head.Next
	}

	return nil
}

//Runtime: 14 ms, faster than 28.12% of Go online submissions for Linked List Cycle II.
//Memory Usage: 5.2 MB, less than 12.07% of Go online submissions for Linked List Cycle II.
