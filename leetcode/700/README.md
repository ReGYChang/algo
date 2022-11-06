# Problem

You are given the `root` of a binary search tree (BST) and an integer `val`.

Find the node in the BST that the node's value equals `val` and return the subtree rooted with that node. If such a node does not exist, return `null`.

## Example 1:

```
Given the tree:
        4
       / \
      2   7
     / \
    1   3

You should return this subtree:
      2
     / \
    1   3
```

- Input: root = [4,2,7,1,3], val = 2
- Output: [2,1,3]

## Example 2:

```
        4
       / \
      2   7
     / \
    1   3
```

- Input: root = [4,2,7,1,3], val = 5
- Output: []
 
## Constraints:

- The number of nodes in the tree is in the range `[1, 5000]`.
- `1 <= Node.val <= 107`
- `root` is a binary search tree.
- 1 <= val <= $10^7$

# Approach 1: Recursive

## Algorithm

利用遞迴完成 binary search

## Implementation

```go
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
```


## Complexity Analysis

- Time complexity: $O(logn)$
  - Where `n` is the number of nodes in the given tree.
  - Runtime: 20 ms, faster than 92.46% of Go online submissions.

- Space complexity: $O(logn)$
  - Where `n` is the number of nodes in the given tree.
  - Memory Usage: 6.9 MB, less than 90.16% of Go online submissions.