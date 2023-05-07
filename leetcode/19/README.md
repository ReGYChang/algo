# Problem

Given the head of a linked list, remove the nth node from the end of the list and return its head.

## Example 1

- Input: head = [1,2,3,4,5], n = 2
- Output: [1,2,3,5]

## Example 2

- Input: head = [1], n = 1
- Output: []

## Example 3

- Input: head = [1,2], n = 1
- Output: [1]

## Constraints

- The number of nodes in the list is sz.
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

# Approach 1:

## Algorithm

1. The function takes two input arguments: a pointer to the head node of the linked list (`head *ListNode`) and an integer n (`n int`).

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
```

2. A new slice `s` of `*ListNode` pointers is created. This slice will store pointers to each node in the linked list.

```go
s := make([]*ListNode, 0)
```

3. The function iterates through the linked list using a `for` loop, adding pointers to each `ListNode` to the slice `s`.

```go
for node := head; node != nil; node = node.Next {
    s = append(s, node)
}
```

4. The n-th node from the end of the list (determined by `len(s) - n`) is set to `nil` in the slice `s`. This signifies that this node should be removed from the linked list.

```go
s[len(s) - n] = nil
```

5. A new empty `ListNode` is created, which serves as the head of the resulting linked list `res`. Another pointer, `t`, is assigned to the same location as `res`. This pointer will be used to traverse and modify the new linked list.

```go
res := &ListNode{}
t := res
```

6. The function iterates through the slice `s` using another `for` loop.

```go
for i := 0; i < len(s); i++ {
```

7. If the current `ListNode` pointer in `s` is not `nil`, it means this node should be included in the new linked list. A new `ListNode` with the same value is created and appended to the `res` linked list. The pointer `t` is then updated to point to the newly added node.

```go
if s[i] != nil {
    t.Next = &ListNode{Val: s[i].Val}
    t = t.Next
}
```

8. If the current `ListNode` pointer in `s` is `nil`, it means this node should be skipped (removed from the linked list). However, if the current node is not the last one in the slice, the next node in the slice should be included in the new linked list. A new `ListNode` with the same value as the next node is created and appended to the `res` linked list. The pointer `t` is then updated to point to the newly added node, and the index `i` is incremented by 1 to skip the next node in the `s` slice, as it has already been processed.

```go
else {
    if i < len(s) - 1 {
        t.Next = &ListNode{Val: s[i+1].Val}
        t = t.Next
        i++
    }
}
```

9. The loop ends, and the new linked list has been constructed without the n-th node from the end.

10. The function returns the new linked list, starting from the first node after the initial empty `ListNode`.

```go
return res.Next
}
```

The overall purpose of the function `removeNthFromEnd` is to create a new singly-linked list without the n-th node from the end of the input linked list. The function achieves this by first iterating through the input list, storing pointers to the nodes in a slice, and then constructing a new list while skipping the node that should be removed.

In summary, the `removeNthFromEnd` function performs the following steps:

1. Initialize an empty slice to store pointers to nodes in the input list.
2. Iterate through the input list and store pointers to its nodes in the slice.
3. Set the n-th node from the end in the slice to `nil`, indicating it should be removed.
4. Initialize a new empty `ListNode` for the resulting linked list and a pointer `t` to construct the new list.
5. Iterate through the slice, appending nodes from the input list to the new list, while skipping the node set to `nil`.
6. Return the new linked list, starting from the first node after the initial empty `ListNode`.

This function provides an approachable solution to the problem of removing the n-th node from the end of a singly-linked list, which can be applied across various programming languages and systems.

## Implementation

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    s := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		s = append(s, node)
	}
	s[len(s)-n] = nil

	res := &ListNode{}
	t := res
	for i := 0; i < len(s); i++ {
		if s[i] != nil {
			t.Next = &ListNode{Val: s[i].Val}
			t = t.Next
		} else {
			if i < len(s) - 1{
				t.Next = &ListNode{Val: s[i+1].Val}
				t = t.Next
				i++
			}
		}
	}

	return res.Next
}
```

## Complexity Analysis

### Time complexity: $O(n)$

1. The code iterates through the linked list using a for loop to populate a slice s with pointers to each ListNode. The number of iterations depends on the length of the linked list, which we can denote as n. This loop contributes O(n) to the time complexity.
   
2. The code then iterates through the slice s using another for loop to construct a new linked list without the n-th node from the end. This loop also iterates n times, contributing O(n) to the time complexity.

Since both loops contribute O(n) to the time complexity, and they are executed sequentially, the overall time complexity of the function is O(2n), which can be simplified to O(n).

> Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.

### Space complexity: $O(n)$

1. The space complexity is affected by the additional memory allocated for the slice s and the new linked list res.

2. The length of the slice s is equal to the length of the input linked list, which is n. Thus, the space complexity contributed by s is O(n).

3. The length of the new linked list res is also equal to L (with one node removed). Therefore, the space complexity contributed by res is O(n).

Since both s and res contribute O(n) to the space complexity, the overall space complexity of the function is O(2n), which can be simplified to O(n).

In summary, the time complexity of the given code is O(n), and the space complexity is O(n). This analysis is applicable across various programming languages and systems, as it focuses on the core algorithm and logic rather than language-specific details.

> Memory Usage: 2.4 MB, less than 5% of Go online submissions for Remove Nth Node From End of List.