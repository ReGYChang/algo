package main

import "algo/structure"

func mergeTwoLists(list1 *structure.ListNode, list2 *structure.ListNode) *structure.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val >= list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
