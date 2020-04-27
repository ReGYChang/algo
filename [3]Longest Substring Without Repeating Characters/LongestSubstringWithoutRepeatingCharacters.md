# [3]Longest Substring Without Repeating Characters

## Explanation
Given a string, find the length of the longest substring without repeating characters.
```
Example: 

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```             
## Tag
> hash-table | two-pointers | string | sliding-window
## Corner Cases
> 遇到重複字元時的處理
```
Input: "" || "a" || "aaaaa" 
```
## Solution
### Approach 1 - HashMap
#### Algo Goal
> 找出最長不重複子字串
#### Processing
- curr: 記錄子字串之起始位子
- subStringMap<子字串字元,字元位置>: 紀錄字元在當前子字串的位置
```
疊代字串，如遇到未記錄的新字元則直接紀錄進subStringMap

否則檢查 curr 是否在當前字元的上個位置之前
* True: 1. 將curr指標移至當前字元上個位置下一格
        2. 紀錄當前字元新位置
* False: 直接記錄當前字元新位置

檢查當前子字串長度
```
#### Code
```JAVA
    public int lengthOfLongestSubstring(String s) {
        int ans = 0, curr = 0, len = s.length();
        String [] stringTokens = s.split("");
        Map<String,Integer> subStringMap = new HashMap<>();

        for(int i = 0; i < len; i++){
            if(!subStringMap.containsKey(stringTokens[i])){
                subStringMap.put(stringTokens[i], i);
            }else{
                if(curr <= subStringMap.get(stringTokens[i])){
                    curr = subStringMap.get(stringTokens[i]) + 1;
                    subStringMap.put(stringTokens[i], i);
                }else{
                    subStringMap.put(stringTokens[i], i);
                }
            }
            if((i - curr + 1) > ans) ans = i - curr + 1;
        }
        return ans;
        
    }
```
#### Analysis
* 987/987 cases passed (19 ms)
* Your runtime beats 20.12 % of java submissions
* Your memory usage beats 5.91 % of java submissions (40.8 MB)

    Time Complexity: O(max(m,n)) 
    
    Space Complexity: O(max(m,n) + 1)
---


