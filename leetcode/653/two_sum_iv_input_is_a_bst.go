package main

import (
	"algo/structure"
)

type TreeNode = structure.TreeNode

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	setMap(root, m)

	for key, _ := range m {
		if count, exist := m[k-key]; exist && !(key == k-key && count < 2) {
			return true
		}
	}
	return false
}

func setMap(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	m[root.Val]++
	setMap(root.Left, m)
	setMap(root.Right, m)
}
