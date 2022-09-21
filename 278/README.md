# Problem

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of
your product fails the quality check. Since each version is developed based on the previous version, all the versions
after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the
following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find
the first bad version. You should minimize the number of calls to the API.

## Example 1:

- Input: n = 5, bad = 4
- Output: 4
- Explanation:
    - call isBadVersion(3) -> false
    - call isBadVersion(5) -> true
    - call isBadVersion(4) -> true
    - Then 4 is the first bad version.

## Example 2:

- Input: n = 1, bad = 1
- Output: 1

## Constraints:

- 1 <= bad <= n <= 231 - 1

# Approach 1: Binary Search

## Algorithm

題目要求於給定正整數中找到指定正整數, 可以利用 `binary search` 的方式從 `median` 開始核實, 核實結果分為兩種情況:

  - 中位數為 `bad version`: 表示大於 `median` 的數字皆為 `bad version`, `median - 1` <= `first bad version` <= `n`
    - 若  `median`  - 1 為 `bad version`: 則 `first bad version` <  `median` ;
    - 若  `median`  - 1 不為 `bad version`: 則表示 `median` 即 `first bad version`, return;
  
- `median` 不為 `bad version`: 表示小於 `median` 的數字皆不為 `bad version`, 則 `first bad version` >  `median` 

## Implementation

```go
func firstBadVersion(n int) int {
	base := 1
	for base <= n {
		mid := (base + n) / 2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			n = mid - 1
		} else {
			base = mid + 1
		}
	}
	return 1
}
```

## Complexity Analysis

- Time complexity: O(log n)

  - For each round will search half the numbers, so it will have log n rounds.
  - Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.

- Space complexity: O(1)

  - Only constant space is used.
  - Memory Usage: 1.8 MB, less than 87.98% of Go online submissions for First Bad Version.