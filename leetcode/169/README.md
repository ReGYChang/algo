# Problem

Given an array `nums` of size `n`, return the majority element.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

## Example 1:

- Input: nums = [3,2,3]
- Output: 3

## Example 2:

- Input: nums = [2,2,1,1,1,2,2]
- Output: 2

## Constraints:

- `n == nums.length`
- `1 <= n <= 5 * 104`
- `-109 <= nums[i] <= 109`

# Approach 1: Hash Table

## Algorithm

利用 hash map 紀錄每過數字出現的頻率, 遍歷 map 找到出現頻率最高的數字返回

## Implementation

```go
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var maxKey, maxVal int
	for k, v := range m {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}

	return maxKey
}
```


## Complexity Analysis

- Time complexity: $O(n)$
  - Where `n` is the number of given array.
  - Runtime: 25 ms, faster than 75.69% of Go online submissions.

- Space complexity: $O(n)$
  - Where `n` is the number of given array.
  - Memory Usage: 6 MB, less than 99.31% of Go online submissions.