- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Example 3](#example-3)
  - [Constraints](#constraints)
- [Approach 1: Recursive](#approach-1-recursive)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

## Example 1

```
1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
               ↓
1 -> 2 -> 3 -> 4 -> 5
```

- Input: head = [1,2,6,3,4,5,6], val = 6
- Output: [1,2,3,4,5]


## Example 2

- Input: head = [], val = 1
- Output: []

## Example 3

- Input: head = [7,7,7,7], val = 7
- Output: []

## Constraints

- The number of nodes in the list is in the range `[0, 104]`.
- `1 <= Node.val <= 50`
- `0 <= val <= 50`

# Approach 1: Recursive

## Algorithm

使用 recursive 的方式返回 `sub-linked list`:

- if 當前 node value 為 `nil`: 返回 `nil`
- if 當前 node value 為需移除的 `val`: 返回剩下未完成 recursive 的 `sub-linked list`
- 返回剩下 `sub-linked list` 即可 

## Implementation

```go
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}

	head.Next = removeElements(head.Next, val)
	return head
}
```

## Complexity Analysis

- Time complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Runtime: 7 ms, faster than 90.19% of Go online submissions for Remove Linked List Elements.

- Space complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Memory Usage: 5 MB, less than 14.07% of Go online submissions for Remove Linked List Elements.
