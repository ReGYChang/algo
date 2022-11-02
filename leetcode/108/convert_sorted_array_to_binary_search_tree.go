package main

import "algo/structure"

type TreeNode = structure.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[0:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return root
}
