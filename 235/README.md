- [Problem](#problem)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Example 3](#example-3)
  - [Constraints](#constraints)
- [Approach 1: Recursion](#approach-1-recursion)
  - [Algorithm](#algorithm)
  - [Implementation](#implementation)
  - [Complexity Analysis](#complexity-analysis)

# Problem

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow `a node to be a descendant of itself`).”

## Example 1

```go
      6
     / \
  2       8
 / \     / \
0   4   7   9
   / \
  3   5
```

- Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
- Output: 6
- Explanation: The LCA of nodes 2 and 8 is 6.

## Example 2

- Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
- Output: 2
- Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

## Example 3

- Input: root = [2,1], p = 2, q = 1
- Output: 2

## Constraints

- The number of nodes in the tree is in the range [2, $10^5$].
- $-10^9$ <= Node.val <= $10^9$
- All Node.val are `unique`.
- p != q
- p and q will exist in the BST.

# Approach 1: Recursion

## Algorithm

遞迴判斷 `p`, `q` 兩節點位於當前節點的左子樹還是右子樹:

- 若位於左子樹則繼續向左尋找最小共同祖先
- 若位於右子樹則繼續向右尋找最小共同祖先
- 否則當前節點則為兩節點最小共同祖先

## Implementation

```go
// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
```

## Complexity Analysis

- Time complexity: O($log{n}$)
  - Where `n` is the height of given tree.
  - Runtime: 41 ms, faster than 58.20% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.

- Space complexity: O($log{n}$)
  - Where `n` is the height of given tree.
  - Memory Usage: 7.2 MB, less than 80.32% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
