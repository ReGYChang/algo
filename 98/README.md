- [Problem](#problem)
	- [Example 1](#example-1)
	- [Example 2](#example-2)
	- [Constraints](#constraints)
- [Approach 1: Stack based iterative](#approach-1-stack-based-iterative)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)
- [Approach 2: Morris Traversal](#approach-2-morris-traversal)
	- [Algorithm](#algorithm-1)
	- [Implementation](#implementation-1)
	- [Complexity Analysis](#complexity-analysis-1)

# Problem

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A `valid BST` is defined as follows:

- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

## Example 1

```go
 2
 | \
 1  3
```

- Input: root = [2,1,3]
- Output: true

## Example 2

- Input: root = [5,1,4,null,null,3,6]
- Output: false
- Explanation: The root node's value is 5 but its right child's value is 4.

## Constraints

- The number of nodes in the tree is in the range `[1, 104]`.
- `-231 <= Node.val <= 231 - 1`

# Approach 1: Stack based iterative

## Algorithm

對於樹中每個節點而言, 需要判斷當前節點是否大於前一個節點, 因此透過 `Inorder traversal` 的方式走訪樹中每個節點並與前一節點比較

使用 `stack` 來完成 `inorder traversal`, 過程如下:

- 創建一個 `slice` 作為 `stack` 使用, 並將 `root` push to stack
- 使用 `previous` 保存前一個節點值
- while `stack` is not empty:
  - pop node from stack
  - push right child of a poped node to stack
  - push left child of a popped node to stack

## Implementation

```go
// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pNode := root
	previous := -2 << 31
	for len(stack) > 0 || pNode != nil {
		if pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Val <= previous {
				return false
			}
			previous = node.Val
			pNode = node.Right
		}
	}
	return true
}
```

## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the number of nodes in the given tree.
  - Runtime: 8 ms, faster than 100% of Go online submissions for Validate Binary Search Tree.

- Space complexity: O(n)
  - The max stack space can be `n` if the tree is left skewed one.
  - Memory Usage: 5.2 MB, less than 85.09% of Go online submissions for Validate Binary Search Tree.

# Approach 2: Morris Traversal

## Algorithm

Morris Traversal 可以做到不使用遞迴或 stack 即可完成 tree traversal, 優化空間複雜度至 `O(1)`

> 需要構建一種機制使得左子樹遍歷完後能返回對應父節點並繼續遍歷幼子樹



## Implementation

```go
// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTMorris(root *TreeNode) bool {
	cur, pre := root, root
	preVal := -2 << 31
	for cur != nil {
		if cur.Left == nil {
			if cur.Val <= preVal {
				return false
			}
			preVal = cur.Val

			cur = cur.Right
		} else {
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				if cur.Val <= preVal {
					return false
				}
				preVal = cur.Val

				pre.Right = nil
				cur = cur.Right
			}
		}
	}
	return true
}
```

## Complexity Analysis

- Time complexity: O(n)
  - Where `n` is the number of nodes in the given tree.
  - Runtime: 7 ms, faster than 73.40% of Go online submissions for Validate Binary Search Tree.

- Space complexity: O(1)
  - Only constant space is used.
  - Memory Usage: 5 MB, less than 96.39% of Go online submissions for Validate Binary Search Tree.
