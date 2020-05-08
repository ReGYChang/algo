# [12]Integer To Roman

## Explanation
Roman numerals are represented by seven different symbols: **I**, **V**, **X**, **L**, **C**, **D** and **M**.

```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example, two is written as **II** in Roman numeral, just two one's added together. Twelve is written as, **XII**, which is simply **X** + **II**. The number twenty seven is written as **XXVII**, which is **XX** + **V** + **II**.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not **IIII**. Instead, the number four is written as **IV**. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as **IX**. There are six instances where subtraction is used:

- **I** can be placed before **V** (5) and **X** (10) to make 4 and 9. 
- **X** can be placed before **L** (50) and **C** (100) to make 40 and 90. 
- **C** can be placed before **D** (500) and **M** (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:
```
Input: 3
Output: "III"
```
Example 2:
```
Input: 4
Output: "IV"
```
Example 3:
```
Input: 9
Output: "IX"
```
Example 4:
```
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```
Example 5:
```
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```
<br>

## Tag
> math | string
<br>

## Corner Cases
> 注意羅馬數字有小到大的數字

Ex : 9 `->`  **IX** `-x->` **XIIII**
<br>

## Solution
### Approach 1 - Recursion
#### Algo Goal
> 遞迴記錄由大到小羅馬數字直到 int < 1
#### Processing
依值域由大到小依次紀錄轉換為相對應羅馬數字

但要注意如**900**或**400**這種例外並額外做處理
#### Code
```JAVA
    public String intToRoman(int num) {
        StringBuilder ans = new StringBuilder();
        int curr = num;
        return convertIntToRoman(ans, curr);
    }

    private String convertIntToRoman(StringBuilder stringBuilder, int curr) {
        StringBuilder s = stringBuilder;
        if(curr >= 1000){
            for(int i = 0; i < curr / 1000; i++){
                s.append('M');
            }
            return convertIntToRoman(s, curr - (curr / 1000) * 1000);
        }else if(curr >= 500){
            if(curr >= 900){
                s.append("CM");
                return convertIntToRoman(s, curr - 900);
            }else{
                s.append('D');
                return convertIntToRoman(s, curr - 500);
            }
        }else if(curr >= 100){
            if(curr >= 400){
                s.append("CD");
                return convertIntToRoman(s, curr - 400);
            }else{
                for(int i = 0; i < curr / 100; i++){
                    s.append('C');
                }
                return convertIntToRoman(s, curr - (curr / 100) * 100);
            }
        }else if(curr >= 50){
            if(curr >= 90){
                s.append("XC");
                return convertIntToRoman(s, curr - 90);
            }else{
                s.append('L');
                return convertIntToRoman(s, curr - 50);
            }
        }else if(curr >= 10){
            if(curr >= 40){
                s.append("XL");
                return convertIntToRoman(s, curr - 40);
            }else{
                for(int i = 0; i < curr / 10; i++){
                    s.append('X');
                }
                return convertIntToRoman(s, curr - (curr / 10) * 10);
            }
        }else if(curr >= 5){
            if(curr >= 9){
                s.append("IX");
                return convertIntToRoman(s, curr - 9);
            }else{
                s.append('V');
                return convertIntToRoman(s, curr - 5);
            }
        }else if(curr >= 1){
            if(curr >= 4){
                s.append("IV");
                return convertIntToRoman(s, curr - 4);
            }else{
                for(int i = 0; i < curr; i++){
                    s.append('I');
                }
            }
        }
        return s.toString();
    }
```
#### Analysis
* 3999/3999 cases passed (3 ms)
* Your runtime beats 100 % of java submissions
* Your memory usage beats 23.75 % of java submissions (38.6 MB)

    Time Complexity: O(1) 
    
    Space Complexity: O(1)

`Note : 由於題目限制範圍 1 to 3999，故複雜度均為O(1)` 

### Approach 2 - Array
#### Algo Goal
> 由大到小依次轉換
#### Processing
將所有可能的羅馬數字與阿拉伯數字組合塞進Array

可以避免處理例外狀況
#### Code
```JAVA
    public String intToRoman(int num) {
        StringBuilder ans = new StringBuilder();
        String[] stringArray = {"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"};
        int[] intArray = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        for(int i = 0; i < 13; i++){
            if(num >= intArray[i]){
                int count = num / intArray[i];
                num %= intArray[i];
                for(int j = 0; j < count; j++) 
                    ans.append(stringArray[i]);
            }
        }
        return ans.toString();
    }
```
#### Analysis
- 3999/3999 cases passed (4 ms)
- Your runtime beats 90.34 % of java submissions
- Your memory usage beats 18.75 % of java submissions (39.4 MB)

    Time Complexity: O(1) 
    
    Space Complexity: O(1)
    
    
