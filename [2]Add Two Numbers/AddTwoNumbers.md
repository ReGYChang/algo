# [2]Add Two Numbers

## Explanation
You are given `two non-empty linked lists` representing two `non-negative` integers. The digits are stored in `reverse` order and each of their nodes contain a `single digit`. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
```
Example: 

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```             
## Tag
> linked-list | math
## Corner Cases
> 注意node相加的進位問題
```
Input: (9 -> 7 -> 8) + (7 -> 9 -> 6)
Output: 6 -> 7 -> 5 -> 1
```
> 注意兩個linked List有不同Node數
```
Input: (1 -> 2) + (4 -> 9 -> 6)
Output: 5 -> 1 -> 7
```
## Solution
### Approach 1 - Iteratively Add Two Nodes
#### Algo Goal
> 對兩Linked List每個Node相加直到兩個Linked List為空為止
#### Processing
- 變數宣告
```
int isCarry: 確認每回合進位值
ListNode currListNode: 當前回合ListNode Pointer
```
- 相加處理
```
將兩ListNode相加並加上上一回合isCarry
每回合isCarry = 相加結果 / 10
存放結果進ListNode並將Pointer往後
```
#### Code
```JAVA
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int isCarry = 0;
        ListNode resultListNode = new ListNode(0);
        ListNode currListNode = resultListNode;
        while (l1 != null || l2 != null) {
            int cellResult = 0;
            int x = (l1 != null) ? l1.val : 0;
            int y = (l2 != null) ? l2.val : 0;

            cellResult = x + y + isCarry;
            isCarry = cellResult / 10;
            cellResult %= 10;

            if (l1 != null)
                l1 = l1.next;
            if (l2 != null)
                l2 = l2.next;

            currListNode.next = new ListNode(cellResult);
            currListNode = currListNode.next;
        }
        
        if (isCarry > 0)
            currListNode.next = new ListNode(isCarry);
        return resultListNode.next;
    }
```
#### Analysis
* 1563/1563 cases passed (1 ms)
* Your runtime beats 100 % of java submissions
* Your memory usage beats 99.69 % of java submissions (39.6 MB)

    Time Complexity: O(max(m,n)) 
    
    Space Complexity: O(max(m,n) + 1)
---


