# Problem

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return -1.

You must write an algorithm with `O(log n)` runtime complexity.

## Example 1:

- Input: nums = [-1,0,3,5,9,12], target = 9
- Output: 4
- Explanation: 9 exists in nums and its index is 4

## Example 2:

- Input: nums = [-1,0,3,5,9,12], target = 2
- Output: -1
- Explanation: 2 does not exist in nums so return -1

## Constraints:

- 1 <= nums.length <= 104
- -104 < nums[i], target < 104
- All the integers in nums are `unique`.
- nums `is sorted in ascending order`.
  
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