- [Problem](#problem)
	- [Example 1](#example-1)
	- [Example 2](#example-2)
	- [Constraints](#constraints)
- [Approach 1: DFS](#approach-1-dfs)
	- [Algorithm](#algorithm)
	- [Implementation](#implementation)
	- [Complexity Analysis](#complexity-analysis)

# Problem

Given an $m*n$ 2D binary grid grid which represents a map of '`1`'s (land) and '`0`'s (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Example 1

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]

Output: 1
```


## Example 2

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]

Output: 3
```

## Constraints

- `m == grid.length`
- `n == grid[i].length`
- 1 <= `m`, `n` <= 300
- `grid[i][j]` is '`0`' or '`1`'.

# Approach 1: DFS

## Algorithm

創建一個 $m*n$ slice `table` 用做標記, `table[m][n]` 表示 `grid[m][n]` 已經遍歷過; 另外使用 `count` 變數紀錄 `land(1)` 數量

遍歷 `grid[m][n]`, 並透過 `DFS recursion` 的方式找到所有局域範圍內的 `land(1)`, 每個 `recursion` 都會上下左右遞迴搜尋未遍歷過的點

一個 `DFS recursion` 結束即代表一個完整的 `land(1)`, 則繼續遍歷下一個未遍歷的點, 直到 `grid` 全部點遍歷結束為止

## Implementation

```go
func numIslands(grid [][]byte) int {
	table := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(table); i++ {
		table[i] = make([]bool, len(grid[i]), len(grid[i]))
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
```

## Complexity Analysis

- Time complexity: $O(m*n)$
  - Where `m == grid.length` and `n == grid[i].length`.
  - Runtime: 3 ms, faster than 90.84% of Go online submissions for Number of Islands.

- Space complexity: $O(m*n)$
  - Where `m == grid.length` and `n == grid[i].length`.
  - Memory Usage: 4 MB, less than 46.68% of Go online submissions for Number of Islands.
