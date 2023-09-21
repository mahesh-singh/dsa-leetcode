# DSA, Leetcode in Golang

### 424. Longest Repeating Character Replacement

Approach #1
1. Given an string, sub string need to find of dynamic length - sliding window
2. l,r = 0, map to track the frequency of the english chars, maxLength to track the max length
3. add char in map or if already exist, increment the count, increase the maxLength
4. find the max frequency char from map and apply following formula  `maxlength - maxFrequency <= k`
5. if `maxlength - maxFrequency <= k` is true, keep repeating above steps
6. if false, reduce the frequency of s[l], move l = l+1  

Approach #2
1. Using binary search


### 121. Best Time to Buy and Sell Stock

Approach #1
1. l & r = 1, maxDiff = 0
2. for `r < len(prices)`
3. if `prices[r] < prices[l]`, l need to shift to new index `l=r`
4. else max of maxDiff and  `prices[r] - prices[l]` 
5. `r = r+1`


### 3. Longest Substring Without Repeating Characters
Approach #1
1. l & r to track the window, set l = 0 , r = 0, run the while loop `r < len(s)`
2. map to track the repeating chars, key=char, value=index
3. maxLength to track the max length
4. loop `r<len(s)`
5. if s[r] not exist in map, take max maxLength and r-l+1
6. else (its seen ), change the `l` to seen index +1
7. add char and index into map


Mistake
Approach #1
1. map to track the repeated char, maxLength to track the max length
2. loop though the array, for each char 
3. if char exists in the map, it means repeating, clear the map, store current char // mistake as it not tracking the index of chars and it will not calculate length correctly  
4. If char not exists in the map, add char in the map, take max of maxLength and len of map
5. return the maxLength

### 33. Search in Rotated Sorted Array

Approach #1

1. In rotated sorted array, if split from mid, at least one side will be increasing order. This indicate rotation happen on other side.
2. find the min element index (while l<r) , this can be done via checking which side is non increasing
3. Once min element index found, run the binay search for for both soid find the target element 

Approach #2
1. Same as above but searching can be happen at the same time
2. while checking increasing order side, we can also check which side target possible can be by checking `target` in between, and change the left or right index accordingly
3. Once l>r condition break, check the target value in `nums[l]==traget` else return -1

Mistakes
1. miss the `target` condition 
2. miss the loop will break and need to check with `nums[l]`


### 75. Search a 2D Matrix

m: number of rows
n: number of column

Approach #1

1. upper bond `i>0, i<m, j>0, j<n` 
2. Start with top right `matrix[0][n-1]` (consider it as mid in binary search)
3. if it is greater then `target`, it can't be in the next row, move to the previous column 
4. if it is less then `target`, move to next row 
5. do this until upper bound reached

Mistake 
1. missing the case for [[1]],1. condition was `j>0;` but it should be `j>=0`

Approach #2

1. Binary search on rows to find probable row index  
2.  if `target` > mid row last element; change change the top
3.  if `target` < mid row fist element; change the bottom
4.  Otherwise you are on the probable row; break the loop   
5. Binary search on the the elements of the row

Complexity  = log(m) + log(n) 

Mistake
1. Edge case where top become greater then bot and while searching for the candidate row and exiting the loop 

Approach #3
Alternative to above approach: convert matrix into array and vise versa 

A n * m matrix converted into an array: matrix[x][y] : a[x * m + y]
An array can be converted into n * m matrix: a[x] : matrix[x / m][x % m]

### 0744 Find Smallest Letter Greater Than Target

1. Assume first element is smallest greater then target
2. Binary search with change  if mid element > `target`, update the assumption to mid element
3. Once mid element == `target`, return the mid+1 element
4. return the latest assumption

Mistake
1. missing the edge case when `mid` reached to end of the array

## Others
-  Find mid element 

```
mid = l + (r-l)/2
```

- Sliding window 
  - Fixed window : sub array or sub string and find sum or length as a max or min 
  - Dynamic window : given condition min diff, max diff, sum, length etc, but find the sub array or sub string 
