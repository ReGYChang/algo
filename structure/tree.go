package structure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = -1 << 63

func Int2TreeNode(integers []int) *TreeNode {
	n := len(integers)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: integers[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && integers[i] != NULL {
			node.Left = &TreeNode{Val: integers[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && integers[i] != NULL {
			node.Right = &TreeNode{Val: integers[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TreeNode2int(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

// Equal return ture if tn == a
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}
