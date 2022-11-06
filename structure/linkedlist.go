package structure

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Int convert List to []int
func List2Int(head *ListNode) []int {
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("length of linked list out of limit %d", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Int2List convert []int to List
func Int2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// Int2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Int2ListWithCycle(nums []int, pos int) *ListNode {
	head := Int2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

func TailOfList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for head.Next != nil {
		head = head.Next
	}
	return head
}
