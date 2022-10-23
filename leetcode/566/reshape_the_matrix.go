package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var m, n = len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	queue := make([]int, m*n)
	idx := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			queue[idx] = mat[i][j]
			idx++
		}
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	idx = 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i][j] = queue[idx]
			idx++
		}
	}

	return res
}

//func matrixReshape(mat [][]int, r int, c int) [][]int {
//	var m, n = len(mat), len(mat[0])
//	if m*n != r*c {
//		return mat
//	}
//
//	res := make([][]int, r, r)
//	for i := 0; i < r; i++ {
//		res[i] = make([]int, c, c)
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if i < r {
//				res[i][j] = mat[i][j]
//			} else {
//				res[i-r][j+m] = mat[i][j]
//			}
//		}
//	}
//
//	return res
//}
