package main

import "algo/structure"

type ListNode = structure.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	ma := make(map[*ListNode]bool)
	for pa != nil {
		ma[pa] = true
		pa = pa.Next
	}

	pb := headB
	for pb != nil {
		if _, exist := ma[pb]; exist {
			return pb
		}
		pb = pb.Next
	}

	return nil
}
