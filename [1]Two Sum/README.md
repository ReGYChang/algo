# [1]Two Sum

## Explanation
> Given an array of integers, return indices of the two numbers such that they add up to a specific target.<br>You may assume that each input would have exactly one solution, and you may not use the same element twice.
```
Example: 

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```             
<br>

## Tag
> Array , HashTable
<br>

## Corner Cases
> 注意重複之元素不可復用
```
nums = [2,7,7,8,8,9], target = 10
```
<br>

## Solution
### Approach 1
#### Algo Goal
> 找出相加結果與Target相同的不重複兩元素
#### Processing
> 用for loop 在兩兩元素間檢查是否符合條件，如遇到相同元素則Continue，直到找到結果為止。
#### Code
```JAVA
    public int[] twoSum(int[] nums, int target) {
        for(int i = 0;i<nums.length;i++){
            for(int j =0;j<nums.length;j++){
                if(i==j){
                    continue;
                }else{
                    if(nums[i]+nums[j]==target){
                        int[] result = {i,j};
                        return result;
                    }
                }
            }
        }
        return null;
    }
```
#### Analysis
