package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
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

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
//Memory Usage: 2 MB, less than 13.91% of Go online submissions for Middle of the Linked List.
