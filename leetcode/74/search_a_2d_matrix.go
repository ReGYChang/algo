package main

func searchMatrix(matrix [][]int, target int) bool {
	var ml, mr = 0, len(matrix) - 1
	for ml != mr {
		mid := (ml + mr) / 2

		if target > matrix[mid][len(matrix[mid])-1] {
			ml = mid + 1
		} else {
			mr = mid
		}
	}

	var nl, nr = 0, len(matrix[ml]) - 1
	for nl != nr {
		mid := (nl + nr) / 2

		if target > matrix[ml][mid] {
			nl = mid + 1
		} else {
			nr = mid
		}
	}

	return matrix[ml][nl] == target
}
