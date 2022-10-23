package main

import "algo/structure"

func levelOrder(root *structure.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*structure.TreeNode{root}
	queueTmp := make([]*structure.TreeNode, 0)
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
			queueTmp = make([]*structure.TreeNode, 0)
			level = make([]int, 0)
			i = -1
		}
	}

	return res
}
