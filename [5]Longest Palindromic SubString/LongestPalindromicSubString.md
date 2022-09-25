# [5]Longest Palindromic SubString

## Explanation
> Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

```
Example: 

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
--------------------------------------
Input: "cbbd"
Output: "bb"

```             
<br>

## Tag
> string | dynamic-programming
<br>

## Corner Cases
> 空陣列 || 連續相同 || 間隔連續相同 
```
Input: "" || Input: "ac" || Input: "cccc" || Input: "bababd" || Input: "cbbd"
```
<br>

## Solution
### Approach 1 - Brute Force
#### Algo Goal
> 找出兩兩元素間滿足Palindromic之最長子字串
#### Processing
> for loop 疊代兩兩元素，並檢查其是否符合Palindromic
```
定義`len`為兩元素間子字串長度

依序檢查`len`長度所有組合並檢查其是否滿足Palindromic特性
```
#### Code
```JAVA
    public String longestPalindrome(String s) {
        String[] tokens = s.split("");  
        String ans = "";
        int len = tokens.length;
        for (int i = 0; i < tokens.length; i++) {
            for(int j = 0; j < tokens.length - len + 1; j++){
                int start = j, end = j + len - 1;
                if(tokens[start].equals(tokens[end])){
                    ans = checkPalindrome(start, end, tokens);
                }else{
                    continue;
                }
                if(!ans.equals("")) return ans;
            }
            len--;
        }
        return ans;
    }
    private String checkPalindrome(int start, int end, String[] tokens) {
        Stack<String> stack = new Stack<>();
        int midVal = (end - start) / 2;
        String curAns = "";
        for(int i = 0, s = start; i <= midVal; i++, s++){
            stack.push(tokens[s]);
            curAns += tokens[s];
        }
        if((end - start + 1) % 2 == 1) {
            stack.pop();
        }
        for(int i = start + midVal + 1; i <= end; i++){
            String token = stack.pop();
            if(token.equals(tokens[i])){
                curAns += token;
            }else{
                return "";
            }
        }
        return curAns;
    }
```
#### Analysis
* Time Limit Exceeded

    Time Complexity: O($n^3$) 
    
    Space Complexity: O(1)
    
---
### Approach 2 - Dynamic Programming
#### Algo Goal
> 找出兩兩元素間滿足Palindromic之最長子字串
#### Processing
> 利用二維陣列紀錄兩兩元素間是否為Palindromic字串,P[i,j] = string.

```
定義 P[i,j] True -> 為子串列 Si .. Sj 符合Palindromic 
            False -> 不符合

P[i,j] = (Si == Sj && P[i + 1][j - 1])
```

定義`len`為兩元素相距長度

如`len <= 2`，則`len 兩端元素相等`可判斷P[i][i + len]符合Palindromic

若`len > 2`，則須滿足`len 兩端元素相等 && dp_array[i + 1][i + len - 1] == true)` 才可判斷P[i][i + len]符合Palindromic

#### Code
```JAVA
    public String longestPalindrome(String s) {
        if(s.length() <= 1)
            return s;
        boolean [][] dp_array = new boolean[1000][1000];
        String ans = s.substring(0,1);
        int sLen = s.length(), maxLen = 1;
        for(int len = 0; len < sLen; len ++){
            for(int i = 0; i < sLen - len; i ++){
                if(s.charAt(i) == s.charAt(i + len) && (len <= 2 || dp_array[i + 1][i + len - 1] == true)){
                    dp_array[i][i + len] = true;
                    if(len + 1  > maxLen) {
                        ans = s.substring(i, i + len + 1);
                        maxLen = len + 1;
                    }
                }
            }
        }
        return ans;
    }
```
#### Analysis
* 103/103 cases passed (99 ms)
* Your runtime beats 28.92 % of java submissions
* Your memory usage beats 6.45 % of java submissions (47.7 MB)

    Time Complexity: O($n^2$) 
    
    Space Complexity: O($n^2$)
    
---
### Approach 3 - Expand Around Center
#### Algo Goal
> 找出兩兩元素間滿足Palindromic之最長子字串
#### Processing
> 以各元素為中心點向兩邊檢查，直到頭尾元素不同為止
```
需考慮字串長度為奇數或偶數的兩種情況　

Example:    "a`b`a"   ||  "a`bb`a"

每次檢查同時檢查兩種狀況並取其字串長度較長者
```
#### Code
```JAVA
    public String longestPalindrome(String s) {
        if(s.equals("") || s.length() < 1) return "";
        int start = 0, end = 0;
        for(int i = 0; i < s.length(); i++){
            int length1 = expandAroundCenter(s, i, i);
            int length2 = expandAroundCenter(s, i, i + 1);
            int lenFinal = Math.max(length1, length2);
            if(lenFinal > end - start + 1){
                start = i - (lenFinal - 1) / 2;
                end = i + lenFinal / 2;
            }
        }
        return s.substring(start , end + 1);
    }
    private int expandAroundCenter(String s, int left, int right) {
       int R = right, L = left;
       while( L >= 0 && R < s.length() && s.charAt(R) == s.charAt(L) ){
           R ++ ;
           L -- ;
       }
       return R - L - 1;
    }
```
#### Analysis
* 103/103 cases passed (22 ms)
* Your runtime beats 90.11 % of java submissions
* Your memory usage beats 52.02 % of java submissions (37.9 MB)

    Time Complexity: O($n^2$) 
    
    Space Complexity: O(1)
