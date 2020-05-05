# [9]Palindrome Number

## Explanation
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
```
Example: 


Input: 121
Output: true


Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.


Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```     
Follow up:

Coud you solve it without converting the integer to a string?
<br>

## Tag
> math 
<br>

## Corner Cases
- negative number
- even / odd
<br>

## Solution
### Approach 1 - Math
#### Algo Goal
> 將int每一位數拆開並判斷是否為Palindromic
#### Processing
- 判斷int是否為負數(若為負數 return false)
- while loop 將int每一位數拆開存入list
- 用Stack先進後出的特性檢查左右半邊位數是否符合Palindromic
#### Code
```JAVA
    public boolean isPalindrome(int x) {
        if(x < 0) return false
        int curr = x, len = 0;
        Stack stack = new Stack<>();
        List<Integer> list = new ArrayList();
        while(curr >= 1){
            list.add(curr % 10);
            curr /= 10;
            len ++;
        }

        int mid = len % 2 == 0 ? len / 2 - 1 : len / 2;

        for(int i = 0; i <= mid; i++)
            stack.push(list.get(i));
        if(len % 2 == 1) stack.pop();
        for(int i = mid + 1; i < len ; i ++)
            if(stack.pop() != list.get(i)) return false;

        return true;
    }
```
#### Analysis
* 11509/11509 cases passed (10 ms)
* Your runtime beats 22.14 % of java submissions
* Your memory usage beats 5.02 % of java submissions (39.1 MB)

    Time Complexity: O(n) 
    
    Space Complexity: O(n)
    
