package main

func numIslands(grid [][]byte) int {
	table := make([][]bool, len(grid))
	for i := 0; i < len(table); i++ {
		table[i] = make([]bool, len(grid[i]))
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !table[i][j] && grid[i][j] == '1' {
				isIsland(i, j, &table, grid)
				count++
			}
		}
	}

	return count
}

func isIsland(gr, gc int, m *[][]bool, grid [][]byte) {
	if gr < 0 || gr == len(grid) || gc < 0 || gc == len(grid[0]) || grid[gr][gc] == '0' || (*m)[gr][gc] {
		return
	}

	(*m)[gr][gc] = true
	isIsland(gr+1, gc, m, grid)
	isIsland(gr-1, gc, m, grid)
	isIsland(gr, gc+1, m, grid)
	isIsland(gr, gc-1, m, grid)
}
