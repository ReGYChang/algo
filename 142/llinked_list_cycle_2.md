# Approach 1: Iteratively

## Algorithm

`Preorder traversal` 是一種 `depth-first search(DFS)` 方法, `DFS` 問題經常會使用到 recursive 的方式來解決

對於每個節點來說:

- 若節點值為 `nil` 則 return `nil`
- 否則遍歷所有子節點依序搜尋, 並返回所有子節點搜尋結果

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

  Where `n` is the number of nodes in the given n-ary tree.

- Space complexity: O(n)

  the max recursive depth can be `n` if the tree is skewed one.