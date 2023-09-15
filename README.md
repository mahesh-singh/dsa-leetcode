# DSA, Leetcode in Golang



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

### 0744 Find Smallest Letter Greater Than Target

1. Assume first element is smallest greater then target
2. Binary search with change  if mid element > `target`, update the assumption to mid element
3. Once mid element == `target`, return the mid+1 element
4. return the latest assumption

Mistake
1. missing the edge case when `mid` reached to end of the array

## Others
1. Find mid element 

```
mid = l + (r-l)/2
```
