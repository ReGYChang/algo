package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	if numRows >= 1 {
		res[0] = make([]int, 1)
		res[0][0] = 1
	}
	if numRows >= 2 {
		res[1] = make([]int, 2)
		res[1][0] = 1
		res[1][1] = 1
	}

	for i := 2; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j > 0 && j < i {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			} else {
				res[i][j] = 1
			}
		}
	}

	return res
}
