# [7]Reverse Integer

## Explanation
> Given a 32-bit signed integer, reverse digits of an integer.
```
Example: 

Input: 123
Output: 321

Input: -123
Output: -321

Input: 120
Output: 21
```     
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

<br>

## Tag
> math
<br>

## Corner Cases
> int overflow 
```
Input: 0 || Input: 1534236469  
```
Note: JAVA int overflow時會跑到最小值，JDK 8 以上可用`Math.multiplyExact()`
捕捉`ArithmeticException`利用try-catch block處理
<br>

## Solution
### Approach 1 - Math
#### Algo Goal
> 自訂義反轉數字
#### Processing
```
如輸入是負數則先將其改為正數

使用while loop 循序除10取餘數並加回ans

最後判斷ans若使用int是否overflow及正負數
```
#### Code
```JAVA
public int reverse(int x) {
        int curr = x >= 0 ? x : -x;
        long ans = 0;
        while(curr > 0){
            ans = ans * 10 + curr % 10;
            curr /= 10;
        }
        if(ans > Integer.MAX_VALUE){
            return 0;
        }else if(x > 0){
            return (int) ans;
        }else{
            return (int) -ans;
        }
    }
```
#### Analysis
* 1032/1032 cases passed (1 ms)
* Your runtime beats 100 % of java submissions
* Your memory usage beats 5.55 % of java submissions (36.5 MB)

    Time Complexity: O(n) 
    
    Space Complexity: O(1)
    
