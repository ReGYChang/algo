package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	queue := []int{sr, sc}
	origin := image[sr][sc]

	image[sr][sc] = color
	m := make([][]bool, len(image), len(image))
	for i := 0; i < len(m); i++ {
		m[i] = make([]bool, len(image[i]), len(image[i]))
	}

	for len(queue) > 0 {
		curSr := queue[0]
		curSc := queue[1]
		queue = queue[2:]

		if curSr > 0 && !m[curSr-1][curSc] && image[curSr-1][curSc] == origin {
			m[curSr-1][curSc] = true
			image[curSr-1][curSc] = color
			queue = append(queue, curSr-1, curSc)
		}
		if curSr < len(image)-1 && !m[curSr+1][curSc] && image[curSr+1][curSc] == origin {
			m[curSr+1][curSc] = true
			image[curSr+1][curSc] = color
			queue = append(queue, curSr+1, curSc)
		}
		if curSc > 0 && !m[curSr][curSc-1] && image[curSr][curSc-1] == origin {
			m[curSr][curSc-1] = true
			image[curSr][curSc-1] = color
			queue = append(queue, curSr, curSc-1)
		}
		if curSc < len(image[curSr])-1 && !m[curSr][curSc+1] && image[curSr][curSc+1] == origin {
			m[curSr][curSc+1] = true
			image[curSr][curSc+1] = color
			queue = append(queue, curSr, curSc+1)
		}
	}
	return image
}
