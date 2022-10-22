package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleHashMap(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*ListNode]bool)
	for head.Next != nil {
		if _, exist := m[head]; !exist {
			m[head] = true
			head = head.Next
		} else {
			return true
		}
	}
	return false
}

func hasCycleTwoPointer(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
