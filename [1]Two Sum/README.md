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
### Approach 1  Brute Force
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
* 29/29 cases passed (82 ms)
* Your runtime beats 8.28 % of java submissions
* Your memory usage beats 5.65 % of java submissions (39.6 MB)

    Time Complexity: O(n^2) 
    
    Space Complexity: O(1)

### Approach 2  Result Record
#### Algo Goal
> 找出符合目標之兩元素
#### Processing
> 利用HashMap紀錄<元素與目標差距,元素位置>，用一個for loop 檢查當前元素是否相同於HashMap<br>
  已存在之Key，若存在則回傳當前陣列位置與目標Key值之Value。
#### Code
```JAVA
    public int[] twoSum(int[] nums, int target) {
        Map<Integer,Integer> recordMap = new HashMap<>();
        for(int i = 0; i < nums.length; i ++){
                if(recordMap.containsKey(nums[i])){
                    return new int[] {recordMap.get(nums[i]), i};
                }
            recordMap.put(target - nums[i], i);
        }
        return null;
    }
```
#### Analysis
* 29/29 cases passed (1 ms)
* Your runtime beats 99.91 % of java submissions
* Your memory usage beats 5.65 % of java submissions (39.9 MB)

    Time Complexity: O(n) 
    
    Space Complexity: O(1)  
