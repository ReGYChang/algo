# Problem

Given the head of a linked list, return the node where the cycle begins. If
there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can
be reached again by continuously following the next pointer. Internally, pos is
used to denote the index of the node that tail's next pointer is connected to (0
-indexed). It is -1 if there is no cycle. Note that pos is not passed as a
parameter.

> Do not modify the linked list.


## Example 1:

- Input: head = [3,2,0,-4], pos = 1
- Output: tail connects to node index 1
- Explanation: There is a cycle in the linked list, where tail connects to the
second node.


## Example 2:

- Input: head = [1,2], pos = 0
- Output: tail connects to node index 0
- Explanation: There is a cycle in the linked list, where tail connects to the
first node.

## Example 3:

- Input: head = [1], pos = -1
- Output: no cycle
- Explanation: There is no cycle in the linked list.

## Constraints:

- The number of the nodes in the list is in the range [0, 10⁴].
- -10⁵ <= Node.val <= 10⁵
- pos is -1 or a valid index in the linked-list.

## Follow up

Can you solve it using O(1) (i.e. constant) memory?

# Approach 1: Iteratively

## Algorithm

透過 `hash table` 來儲存每個經過的 node, 並隨時檢查當前 node 是否已存在 `hash table` 中, 若存在則代表 linked list 有 cycle, 而當前 node 即為 cycle 的起點

## Implementation

```go
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
```

## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the number of nodes in the given linked list.
  - Runtime: 14 ms, faster than 28.12% of Go online submissions for Linked List Cycle II.

- Space complexity: O(n)
  - The max iterations can be all the number of nodes in the given linked list.
  - Memory Usage: 5.2 MB, less than 12.07% of Go online submissions for Linked List Cycle II.