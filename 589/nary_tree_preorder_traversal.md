- [Problem](#problem)
	- [Example 1](#example-1)
	- [Example 2](#example-2)
	- [Constraints](#constraints)
	- [Follow up](#follow-up)
- [Approach 1: Recursive](#approach-1-recursive)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)
- [Approach 2: Iterative](#approach-2-iterative)
	- [Algorithm](#algorithm-1)
	- [Implementation](#implementation-1)
	- [Complexity Analysis](#complexity-analysis-1)

# Problem

Given the root of an n-ary tree, return the preorder traversal of its nodes'
values.

Nary-Tree input serialization is represented in their level order traversal.
Each group of children is separated by the null value (See examples)


## Example 1

- Input: root = [1,null,3,2,4,null,5,6]
- Output: [1,3,5,6,2,4]

## Example 2

- Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,
null,12,null,13,null,null,14]
- Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

## Constraints

- The number of nodes in the tree is in the range [0, 10⁴].
- 0 <= Node.val <= 10⁴
- The height of the n-ary tree is less than or equal to 1000.

## Follow up

Recursive solution is trivial, could you do it iteratively?

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
  - The max recursive depth can be `n` if the tree is skewed one.
  - Memory Usage: 6.2 MB, less than 34.38% of Go online submissions for N-ary Tree Preorder Traversal.

# Approach 2: Iterative

## Algorithm

`Preorder traversal` 同樣也可以用 iterative 的方式解決, 遍歷 nodes 的過程中需**優先處理最前面的 childNodes 及其 children**

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
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for N-ary Tree Preorder Traversal.

- Space complexity: O(n)
  - The max recursive depth can be `n` if the tree is skewed one.
  - Memory Usage: 4.1 MB, less than 79.22% of Go online submissions for N-ary Tree Preorder Traversal.