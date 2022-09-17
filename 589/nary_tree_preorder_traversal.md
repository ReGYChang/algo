# Approach 1: Recursive

## Algorithm

`Preorder traversal` 是一種 `depth-first search(DFS)` 方法, `DFS` 問題經常會使用到 recursive 的方式來解決

對於每個節點來說:

- 若節點值為 `nil` 則 return `nil`
- 否則遍歷所有子節點依序搜尋, 並返回所有子節點搜尋結果

## Implementation

```go
// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	s := []int{root.Val}
	for i := 0; i < len(root.Children); i++ {
		tmp := preorder(root.Children[i])
		if tmp != nil {
			s = append(s, tmp...)
		}
	}

	return s
}
```


## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the number of nodes in the given n-ary tree.
  - Runtime: 3 ms, faster than 75.23% of Go online submissions for N-ary Tree Preorder Traversal.

- Space complexity: O(n)
  - the max recursive depth can be `n` if the tree is skewed one.
  - Memory Usage: 6.2 MB, less than 34.38% of Go online submissions for N-ary Tree Preorder Traversal.

