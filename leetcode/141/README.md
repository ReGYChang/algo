- [Problem](#problem)
	- [Example 1](#example-1)
	- [Example 2](#example-2)
	- [Example 3](#example-3)
	- [Constraints](#constraints)
- [Approach 1: Hash Map](#approach-1-hash-map)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)
- [Approach 2: Two Pointer](#approach-2-two-pointer)
	- [Algorithm](#algorithm-1)
	- [Implementation](#implementation-1)
	- [Complexity Analysis](#complexity-analysis-1)

# Problem

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` if there is a cycle in the linked list. Otherwise, return false.

## Example 1

```
3 -> 2 -> 0 -> 4
     ↑         ↓
     -----------
```

- Input: head = [3,2,0,-4], pos = 1
- Output: true
- Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).


## Example 2

```
1 -> 2
↑    ↓
------
```

- Input: head = [1,2], pos = 0
- Output: true
- Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

## Example 3

```
1
```

- Input: head = [1], pos = -1
- Output: false
- Explanation: There is no cycle in the linked list.

## Constraints

- The number of the nodes in the list is in the range `[0, 104]`.
- `105 <= Node.val <= 105`
- `pos` is `-1` or a **`valid index`** in the linked-list.

# Approach 1: Hash Map

## Algorithm

使用 hashmap `map[*ListNode]bool` 紀錄 node pointer, 若重複遇到相同 node pointer 則表示 linked list 存在 cycle

## Implementation

```go
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
```

## Complexity Analysis

- Time complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Runtime: 18 ms, faster than 31.96% of Go online submissions for Linked List Cycle.

- Space complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Memory Usage: 6.4 MB, less than 16.93% of Go online submissions for Linked List Cycle.

# Approach 2: Two Pointer

## Algorithm

使用快慢指針, 快指針每次移動兩個 node, 慢指針一次移動一個 node, 若 linked list 存在 cycle 則快指針定會與慢指針交會

## Implementation

```go
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
```

## Complexity Analysis

- Time complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Runtime: 0 ms, faster than 100.00% of Go online submissions for Linked List Cycle.

- Space complexity: O(n)
    - Where `n` is the number of nodes in the given linked list.
    - Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Linked List Cycle.
