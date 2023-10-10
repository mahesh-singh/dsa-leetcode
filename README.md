# DSA, Leetcode in Golang

### 9. Remove Nth Node From End of List

Approach #1: reverse, delete, reverse
1. Reverse the lined list
2. Reach to the node just before that, and use tmp variable to store the next.next node
3. set node.next to tmp variable 

Approach #2: Two pointer 
1. Take a dummy node and add list to that (or use the prev pointer, start will null)
2. First pointer at dummy node, second pointer at n distance further from **head** (not form dummy node) , means n+1 node. means offset between these pointer must be equal to n
3. Move both the pointers one step at a time
4. Once second, pointer reach to the null, it means first pointer is just before the node which need to be removed 
5. remove the node by setting   

Mistake
1. Not assigning dummy.next to head. Confused with merged two list where dummy list start with dummy.next= null as add item one by one.

----

### 143. Reorder List
1. Find middle by using slow and fast pointer. Slow.next will be the head of second list
2. Store Slow.next  and **set slow.next to null** to split the fist and second list
3. Reverse the second list, **prev pointer become the new head** 
4. merge these two list until second list become null because in case of odd length, second list become shorter
   1. to merge, first store the next node of fist and second in temp
   2. set fist.next to second
   3. second.next to temp1
   4. move the fist and second list by setting with temp

----

### 141. Linked List Cycle

Approach #1: Brute force
1. Use hash table and store address of the node/value as a key
2. if found again, its cyclic 


Approach #2 : slow and fast pointer
1. Take a two pointer `slow`,  `fast` and set to list 
2. loop through until list == null && slow and fast pointer is null
   1. move `slow` to next list item and `fast` to next of next
   2. compare the `slow` and  `fast`, if equal, return true
3. return false 

Mistake 
1. Making mistake by not checking slowPointer.Next != nil && fastPointer.Next != nil && fastPointer.Next.Next != nil


----

### 21. Merge Two Sorted Lists

Approach #1
1. Note: list is sorted, unequal list
2. Take a dummy list with {Val:-1, Next: null}. 
3. Assign it to tail //trick, this will help in moving tail in the loop
4. Loop both the list until one of them reach to null
5. Compare the value and the smaller one should be added in the tail and tail need to move to next
6. Once loop end, check the not null list and append on the tail
7. return dummy.next //trick: not tail as tail now point to last node

Mistakes
1. Not using tail, this make dummy list reach to last node so not sure how to return the new list
2. `dummy := &ListNode{Val: -1, Next: nil}` 

-----

### 206. Reverse Linked List

Approach #1 : iterative 
1. have three pointers, prev, curr and next 
2. loop through until curr == null // `making mistake of curr.next == null while coding`
   1. set next = curr.next //otherwise you will lose the next ref
   2. set curr.next = prev
   3. set prev = curr
   4. set curr = next

Mistake 
1. `var prev *ListNode = nil`   

----

### 424. Longest Repeating Character Replacement

Approach #1
1. Given an string, sub string need to find of dynamic length - sliding window
2. l,r = 0, map to track the frequency of the english chars, maxLength to track the max length
3. add char in map or if already exist, increment the count, increase the maxLength
4. find the max frequency char from map and apply following formula  `maxlength - maxFrequency <= k` (*it is not same as `maxlength - maxFrequency -k < 0`. It will lead to *)
5. if `maxlength - maxFrequency <= k` is true, keep repeating above steps
6. if false, reduce the frequency of s[l], move l = l+1  

Approach #2
1. Using binary search


-----

### 121. Best Time to Buy and Sell Stock

Approach #1
1. l & r = 1, maxDiff = 0
2. for `r < len(prices)`
3. if `prices[r] < prices[l]`, l need to shift to new index `l=r`
4. else max of maxDiff and  `prices[r] - prices[l]` 
5. `r = r+1`


------

### 3. Longest Substring Without Repeating Characters
Approach #1
1. l & r to track the window, set l = 0 , r = 0, run the while loop `r < len(s)`
2. map to track chars index, key=char, value=index & maxLength to track the max length
3. loop `r<len(s)`
   1. if s[r] not exist in map, take max of maxLength and r-l+1
   2. else (its seen ), and seen index pos is `>=l` (it means seen already outside of current window on the left side)  change the `l` to seen-index +1 
   3. update char index to `r`

----

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

-----

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
5.  if `top > bottom` means, not found. (if value greater-then last element of the 2D array or less-then fist element of the array)  
6. Binary search on the the elements of the row

Complexity  = log(m) + log(n) 

Mistake
1. Edge case where top become greater then bot and while searching for the candidate row and exiting the loop 

Approach #3
Alternative to above approach: convert matrix into array and vise versa 

A n * m matrix converted into an array: matrix[x][y] : a[x * m + y]
An array can be converted into n * m matrix: a[x] : matrix[x / m][x % m]

------

### 0744 Find Smallest Letter Greater Than Target

1. Assume first element is smallest greater then target
2. Binary search with change  if mid element > `target`, update the assumption to mid element
3. Once mid element == `target`, return the mid+1 element
4. return the latest assumption

Mistake
1. missing the edge case when `mid` reached to end of the array

-----
## Others
-  Find mid element 

```
mid = l + (r-l)/2
```

----

### Sliding window 
  - Fixed window : sub array or sub string and find sum or length as a max or min 
  - Dynamic window : given condition min diff, max diff, sum, length etc, but find the sub array or sub string 

----
### Sliding window vs two pointer 
- Two pointer usually compare two values at the two pointers instead of all the element between the pointers
- Sliding window for sub-array, two pointers a way to process pairs  