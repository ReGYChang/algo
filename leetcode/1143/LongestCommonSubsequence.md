# [1143]Longest Common Subsequence

## Explanation
Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.

 

If there is no common subsequence, return 0.
```
Example 1: 

Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
```             
## Tag
> dynamic-programming
## Corner Cases
> 兩字串長度不一 || 子序列相反
```
Input: text1 = "abcde", text2 = "ace" 

Input: text1 = "adefb", text2 = "bcgan" 
```
## Solution
### Approach 1 - Dynamic-Programming
#### Algo Goal
> 找到兩字串的LCS(Longest Common Subsequence)
```JAVA
dp_array[i][j] = char_array1[i - 1] == char_array2[j - 1] ? 
                    dp_array[i - 1][j - 1] + 1 : Math.max(dp_array[i][j - 1], dp_array[i - 1][j]);
```
#### Processing
定義dp[i][j] 為text1前i個字元與text2前j個字元的LCS

當掃描到text1第i個字元與text2第j個字元時:<br>

`如兩字元相等` -> 可判斷LCS為`text1前(i - 1)個字元與text2前(j - 1)的LCS + 1`; <br>

`如兩字元不相等` -> 則需退一步取`text1前i個字元與text2前(j - 1)個字元的LCS`  **&&**  `text1前(i - 1)個字元與text2前j個字元`的LCS之最大值
#### Code
```JAVA
    public int longestCommonSubsequence(String text1, String text2) {
        int[][] dp_array = new int[1000][1000];
        int len1 = text1.length(), len2 = text2.length();
        char[] char_array1 = text1.toCharArray(), char_array2 = text2.toCharArray();

        for(int i = 1; i <= len1; i++){
            for(int j = 1; j <= len2; j++){
                if(char_array1[i - 1] == char_array2[j - 1]){
                    dp_array[i][j] = dp_array[i - 1][j - 1] + 1;
                }else{
                    dp_array[i][j] = Math.max(dp_array[i][j - 1], dp_array[i - 1][j]);
                }
            }
        }
        return dp_array[len1][len2];
    }
```
#### Analysis
* 37/37 cases passed (17 ms)
* Your runtime beats 45.08 % of java submissions
* Your memory usage beats 100 % of java submissions (49.8 MB)

    Time Complexity: $O(n^2)$ 
    
    Space Complexity: $O(n^2)$
---


