# DSA, Leetcode in Golang

## Binary Search

### 33. Search in Rotated Sorted Array
Apprach one

1. In rotated dorted array, if split from mid, at least one side will be increasing order. This indicate rotation happen on other side.
2. find the min element index (while l<r) , this can be done via cheking which side is non increasing
3. Once min element index found, run the binay search for for both soid find the target element 

Approach two
1. Same as above but searching can be happen at the same time
2. while checking increasing order side, we can also check which side target possiblly can be, and change the left or right index accordingly


## Others
1. Find mid element 

```
mid = l + (r-l)/2
```
