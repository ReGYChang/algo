# [10]Regular Expression Matching

## Explanation
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
```
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```
The matching should cover the entire input string (not partial).

Note:

- s could be empty and contains only lowercase letters a-z.
- p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example 1:
```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```
Example 2:
```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```
Example 3:
```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```
Example 4:
```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```
Example 5:
```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```
<br>

## Tag
> string | dynamic-programming | backtracking
<br>

## Corner Cases
- Pattern中 `*` 元素判斷
    - 前一元素 * 0
    - 前一元素 * 1
    - 前一元素 * n
<br>

## Solution
### Approach 1 - Dynamic Programming
#### Algo Goal
> 通過二維陣列dp[i,j]紀錄p[0..j]是否Matching s[0..i]
#### Processing
題目要求判斷兩字串s,p是否相符，利用dynamic programming的方式記錄中間狀態

p字串中元素可分為兩類 : 
- a-z character or '.'
- '*'

其中'*'的狀態又可分兩類 :
- zero preceding element
- one to more preceding element

故每回合切出三種狀態:
- matchSingle : 當前判斷元素為`a-z chars` 或 `.`
> dp[i][j] = dp[i - 1][j - 1]
- matchZero : 當前元素為 `*` 並且將前一元素視為空
> dp[i][j] = dp[i][j - 2]
- matchMore : 當前元素為 `*` 並且判斷是否與前一元素相同
> dp[i][j] = dp[i - 1][j] && p前一元素 == 當前s元素
```
當滿足三種狀態任一時 dp[i][j] = true;
否則 dp[i][j] = false;
```
#### Code
```JAVA
public boolean isMatch(String s, String p) {
        int sLen = s.length(), pLen = p.length();
        boolean[][] dp = new boolean[sLen + 1][pLen + 1];
        char[] charP = p.toCharArray();
        char[] charS = s.toCharArray();
        dp[0][0] = true;
        for(int i = 0; i < sLen + 1; i++){
            for(int j = 1; j < pLen + 1; j++){
                boolean matchSingle = i > 0 && dp[i - 1][j - 1] && (charS[i - 1] == charP[j - 1] || charP[j - 1] == '.') ? true : false;
                boolean matchZero = j > 1 && charP[j - 1] == '*' && dp[i][j - 2] ? true : false;
                boolean matchMore = j > 1 && i > 0 && charP[j - 1] == '*' && ((dp[i - 1][j] && (charS[i-1] == charP[j-2] || charP[j-2] == '.'))) ? true : false;
                dp[i][j] = matchSingle || matchZero || matchMore ? true : false;
            }
        }

        return dp[sLen][pLen];
    }
```
#### Analysis
* 447/447 cases passed (3 ms)
* Your runtime beats 82.06 % of java submissions
* Your memory usage beats 47.47 % of java submissions (38.2 MB)

    Time Complexity: O(n*m) 
    
    Space Complexity: O(n*m)
    
    
