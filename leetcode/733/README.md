# Problem

Given an `m x n` 2D binary grid grid which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



## Example 1:
```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]

Output: 1
```

## Example 2:

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]

Output: 3
```

## Constraints:

- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] is '0' or '1'.
  
# Approach 1: Binary Search

## Algorithm

題目需要於給定 int slice 中找到指定 int 並返回, 可使用 `binary search` 方式查找 `median`, 若目標 int 大於 `median` 則向 slice 右方查找; 若小於則向 slice 左方查找

> `median` 可透過左右邊界算出

## Implementation

```go
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		n := nums[mid]

		if n > target {
			right = mid - 1
		} else if n < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
```

## Complexity Analysis

- Time complexity: O(log n)

  - For each round will search half the numbers, so it will have log n rounds.
  - Runtime: 33 ms, faster than 91.93% of Go online submissions for Binary Search.

- Space complexity: O(1)

  - Only constant space is used.
  - Memory Usage: 6.6 MB, less than 94.44% of Go online submissions for Binary Search.