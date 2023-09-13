# DSA, Leetcode in Golang

## Binary Search

### 33. Search in Rotated Sorted Array

Approach #1

1. In rotated sorted array, if split from mid, at least one side will be increasing order. This indicate rotation happen on other side.
2. find the min element index (while l<r) , this can be done via cheking which side is non increasing
3. Once min element index found, run the binay search for for both soid find the target element 

Approach #2
1. Same as above but searching can be happen at the same time
2. while checking increasing order side, we can also check which side target possiblly can be, and change the left or right index accordingly

### 75. Search a 2D Matrix

Approach #1




## Others
1. Find mid element 

```
mid = l + (r-l)/2
```
