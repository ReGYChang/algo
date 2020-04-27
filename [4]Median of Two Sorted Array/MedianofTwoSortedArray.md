# [4]Median of Two Sorted Arrays

## Explanation
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
```
Example 1: 
nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```             
## Tag
> array | binary-search | divide-and-conquer
## Corner Cases
> 單邊陣列為空 || 兩邊陣列數量不一樣
```
nums1 = [1, 2]
nums2 = []

--------------
nums1 = [1, 2, 3]
nums2 = [4, 5]
```
## Solution
### Approach 1 - Binary-Search
#### Algo Goal
> 找出混合陣列第(m + n + 1) / 2位數 與 (m + n + 2) / 2位數
```
不論(m + n)為奇偶數，其中位數均為 (m + n + 1) / 2 加 (m + n + 2) / 2 除2
```
#### Processing
- 既然時間複雜度規定在O(log(m + n))，應放棄合併陣列重新排序的想法
- 定義findKth()遞迴找出第kth的數
定義 midVal_1,midVal_2 分別記錄 nums1,nums2 第(k / 2)th的數

希望在混合陣列中找到第Kth的數，先在兩陣列中分別找到`第(K / 2)th`數<br><br>
如果`midVal_1 小於 midVal_2`，表示混合陣列中第Kth數必然不在num1中，則可淘汰num1前(k / 2)個數;反之淘汰num2前(k / 2)個數，直到找到混合陣列第Kth的數。

定義 i,j 分別記錄 nums1,nums2 起始位置，接著處理一些CASES

當某一陣列`起始位置大於等於該陣列長度`時，表示該陣列數字已全部被淘汰。
混合陣列第Kth數可直接判斷為另一陣列 (j + k - 1)th數

當要找第一個數時，直接回傳兩陣列最小的第一個數

判斷num1 / num2中是否有包含第(k / 2)的數，若不包含則無法淘汰該陣列，則將
midVal_1 / midVal_2設為最大值防止被淘汰


#### Code
```JAVA
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length, n = nums2.length,
                left = (m + n + 1) / 2, right = (m + n + 2) / 2;

        return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / 2.0;
    }

    public static int findKth(int[] nums1, int i, int[] nums2, int j, int k) {
        if (i >= nums1.length) return nums2[j + k - 1];
        if (j >= nums2.length) return nums1[i + k - 1];
        if (k == 1) return Math.min(nums1[i], nums2[j]);
        int midVal_1 = (i + k / 2 - 1 < nums1.length) ? nums1[i + k / 2 - 1] : Integer.MAX_VALUE;
        int midVal_2 = (j + k / 2 - 1 < nums2.length) ? nums2[j + k / 2 - 1] : Integer.MAX_VALUE;
        if(midVal_1 < midVal_2){
            return findKth(nums1, i + k / 2, nums2, j, k - k / 2);
        }else{
            return findKth(nums1, i, nums2, j + k / 2, k - k / 2);
        }
    }
```
#### Analysis
* 2085/2085 cases passed (2 ms)
* Your runtime beats 99.77 % of java submissions
* Your memory usage beats 100 % of java submissions (40.8 MB)

    Time Complexity: O(log(m + n)) 
    
    Space Complexity: O(1)
---


