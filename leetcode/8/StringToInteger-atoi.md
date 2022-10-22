# [8]String To Integer (atoi)

## Explanation
Implement `atoi` which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
```
Example: 

Input: "42"
Output: 42


Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.Then take as many numerical digits as possible, which gets 42.


Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.


Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical digit or a +/- sign. Therefore no valid conversion could be performed.


Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
```     
<br>

## Tag
> math | string
<br>

## Corner Cases
- null or empty string
- +/- sign
- if overflow or underflow
<br>

## Solution
### Approach 1 - Math
#### Algo Goal
> 找出符合條件之字串並轉換成Integer
#### Processing
- 先用String 的trim()方法將字串前後whitespace去除
- 判斷字串第一個字元為`'+' | '-' | 空字元 | int字元`
- 利用字元的ASCII編碼來判斷並轉換成int
- 判斷正負號及是否溢位並回傳答案
#### Code
```JAVA
public int myAtoi(String str) {
        double ans = 0;
        boolean sign = true;
        int pointer = 0;
        String inputString = str.trim();
        char [] inputChar = inputString.toCharArray();
        
        if(inputChar.length > 0 && inputChar[0] == '-' ){
            sign = false;
            pointer ++;
        }else if(inputChar.length > 0 && inputChar[0] == '+' ){
            pointer ++;
        }else if(inputChar.length <= 0 || (!(inputChar[0] >= '0' && inputChar[0] <= '9' ))){
            return 0;
        }
        while(pointer < inputChar.length && inputChar[pointer] >= '0' && inputChar[pointer] <= '9'){
            ans = ans * 10 + (inputChar[pointer] - '0');
            pointer ++;
        }
        if(sign == false) ans = -ans;
        if(ans > Integer.MAX_VALUE){
            return Integer.MAX_VALUE;
        }else if(ans < Integer.MIN_VALUE){
            return Integer.MIN_VALUE;
        }else{
            return (int) ans;
        }
    }
```
#### Analysis
* 1079/1079 cases passed (2 ms)
* Your runtime beats 79.95 % of java submissions
* Your memory usage beats 5.59 % of java submissions (39.7 MB)

    Time Complexity: O(n) 
    
    Space Complexity: O(n)
    
