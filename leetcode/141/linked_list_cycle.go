package main

import "algo/structure"

func hasCycleHashMap(head *structure.ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*structure.ListNode]bool)
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

func hasCycleTwoPointer(head *structure.ListNode) bool {
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
