# [13]Roman to Integer

## Explanation
Roman numerals are represented by seven different symbols: `I`, `V`, `X`, `L`, `C`, `D` and `M`.

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

For example, two is written as `II` in Roman numeral, just two one's added together. Twelve is written as, `XII`, which is simply `X` + `II`. The number twenty seven is written as `XXVII`, which is `XX` + `V` + `II`.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not `IIII`. Instead, the number four is written as `IV`. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as `IX`. There are six instances where subtraction is used:

- `I` can be placed before `V` (5) and `X` (10) to make 4 and 9. 
- `X` can be placed before `L` (50) and `C` (100) to make 40 and 90. 
- `C` can be placed before `D` (500) and `M` (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:
```
Input: "III"
Output: 3
```

Example 2:
```
Input: "IV"
Output: 4
```

Example 3:
```
Input: "IX"
Output: 9
```

Example 4:
```
Input: "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
```

Example 5:
```
Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```
<br>

## Tag
> math | string
<br>

## Corner Cases
> 注意羅馬字串比前一格大的字元

<br>

## Solution
### Approach 1 - Array
#### Algo Goal
> 掃描字串並將羅馬字元轉換成阿拉伯數字
#### Processing
首先要想辦法讓羅馬字元所代表的阿拉伯數字作對應

直覺會想到使用HashMap用KV的方式記錄

但HashMap用在這裡太笨重了

還是希望能使用Array去完成紀錄對應的動作

這裡使用一個小trick: 利用ASCII當作Index去紀錄數字

優點是速度較快 ; 缺點是浪費空間較大

每回合執行:
- 如果當前char > 前一格char : 表示這格數字應等於當前字元代表的數字 - 前一格字元所代表的數字
- 否則答案加上當前字元代表的數字

#### Code
```JAVA
    public int romanToInt(String s) {
        char[] sChar = s.toCharArray();
        int len = s.length(), ans = 0;
        int [] romanToInt = new int[128];
        romanToInt['I']=1;
        romanToInt['V']=5;
        romanToInt['X']=10;
        romanToInt['L']=50;
        romanToInt['C']=100;
        romanToInt['D']=500;
        romanToInt['M']=1000;

        ans = romanToInt[sChar[0]];
        for(int i = 1; i < len; i++){
          if(romanToInt[sChar[i]] > romanToInt[sChar[i - 1]]){
            ans -= romanToInt[sChar[i - 1]];
            ans += romanToInt[sChar[i]] - romanToInt[sChar[i - 1]];
          }else{
            ans += romanToInt[sChar[i]];
          }
        }
        return ans;
    }
```
#### Analysis
- 3999/3999 cases passed (3 ms)
- Your runtime beats 100 % of java submissions
- Your memory usage beats 5.48 % of java submissions (40.3 MB)

    Time Complexity: O(n) 
    
    Space Complexity: O(1)



