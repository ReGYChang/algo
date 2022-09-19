- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Example 3](#example-3)
  - [Constraints](#constraints)
- [Approach 1: Iteratively](#approach-1-iteratively)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Example 1

```go
 3
 | \
 9  20
    | \
    15 7
```

- Input: root = [3,9,20,null,null,15,7]
- Output: [[3],[9,20],[15,7]]

## Example 2

- Input: root = [1]
- Output: [[1]]

## Example 3

- Input: root = []
- Output: []

## Constraints

- The number of nodes in the tree is in the range [0, 2000].
- -1000 <= Node.val <= 1000

# Approach 1: Iteratively

## Algorithm

此題想法很單純, 透過 Queue FIFO 的特性保存每一層的 node, 遍歷 Queue 同時也記錄下一層 nodes 於 Queue 中, 直到該層最後一個 node 為止; 若有下一層則繼續遍歷, 否則退出並返回結果

## Implementation

```go
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	queueTmp := make([]*TreeNode, 0)
	res := make([][]int, 0)
	level := make([]int, 0)

	for i := 0; i < len(queue); i++ {
		level = append(level, queue[i].Val)

		if queue[i].Left != nil {
			queueTmp = append(queueTmp, queue[i].Left)
		}
		if queue[i].Right != nil {
			queueTmp = append(queueTmp, queue[i].Right)
		}

		if i == len(queue)-1 {
			res = append(res, level)
			queue = queueTmp
			queueTmp = make([]*TreeNode, 0)
			level = make([]int, 0)
			i = -1
		}
	}

	return res
}
```

## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the number of nodes in the given tree.
  - Runtime: 2 ms, faster than 67.53% of Go online submissions for Binary Tree Level Order Traversal.

- Space complexity: O(n)
  - Where `n` is the number of nodes in the given tree.
  - Memory Usage: 2.7 MB, less than 89.92% of Go online submissions for Binary Tree Level Order Traversal.