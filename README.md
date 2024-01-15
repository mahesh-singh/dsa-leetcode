# DSA, Leetcode in Golang

## Pattern 
https://sebinsua.com/algorithmic-bathwater
https://sebinsua.com/assets/posts/algorithmic-bathwater/algorithmic-bathwater-1.svg

# 210. Course Schedule - II
Intuition: Same as course schedule (207), instead of returning the true/false, need to return order of the course schedule via first arranging the course directed graph in adjacency list and then applying topological sort.

Approach #1: Topological sort
1. Convert into adjacency list course -> list of prerequisite
2. visited for tracking the course schedule, visiting for cycle in the graph
3. dfs 
   1. if in visited, return True //already processed
   2. if in visiting, return False // cycle detected 
   3. visiting  = True // on the path to find course without prerequisite
   4. for adj of curr 
      1. call dfs, if dfs return False, return False
   5. set visiting to false // path complete and cycle not detected 
   6. append in visited //
   7. return True
4. for all the course 
   1. call dfs, if return False, return []
5. return visited

------


### 981. Time Based Key-Value Store
Intuition: Key value store will standard map. Time based values, each key will store list of value along with timestamp. Timestamp are sorted by default which means binary search can apply for target timestamp

Approach #1: Binary search
1. take a map where key will string and value will be array of array
2. Set
   1. if key not exist, add key with empty
   2. append the value with timestamp for the key
3. Get
   1. if key not exist return ""
   2. if key exist, get the values which is array of value and timestamp
   3. apply binary search 
      1. if mid less then eq to target
         1. set  return value to values[mid].value //`Trick to return the nearest less then equal`
         2. set l = mid + 1
      2. else
         1. set r = mid - 1
   4. return the return value

-----

### 269 Alien Dictionary
Intuition: Need to return word order. 
1. Lexicographically sorting work based on first different character i.e. `apple`, `aps`; sorting will happen based on `p` and `s` as first two character `ap` is common. 
2.  Build a adjacency list where key will be characters and adjacency list will be characters followed by key according to Lexicographically sorting. Apply topological sort on graph to get the ordering

Approach #1: Topological sort via dfs
1. map to store the ordering
2. build the adjacency list
   1. for i=0 to len(words)-1
      1. get min length of word[i] and word[i+1]
      2. if word1 less then word2 and word1[:min-length] = word2[:min-length] return blank // `not a valid ordering`
      3. for j=0 to min-length
         1. if char of both the words not match
            1. char1 will be key and char2 will be added in adjacency list 
            2. break //`because there is no guarantee that after first non matching char, characters will be in sorted order `
3. for each char key in adjacency list call dfs
   1. if not dfs
      1. return ""
   2. dfs
      1. if in visited return true
      2. if in visiting return false //cycle detected
      3. set visiting true
      4. for each adj of current character 
         1. if not dfs
            1. return false
      5. append in visited 
      6. set visiting false
      7. return true
4. return a reverse string join via visited


reference: https://www.youtube.com/watch?v=IIgisEKjCKA&t=1054s

-----

### 261: Graph Valid Tree (find it on lint code)
Intuition:  Graph is a tree if its is acyclic and connected. Connected means visited node == no of node. Acyclic can be check via DFS (BFS) to check visited with previous node or Union-find by tracking parents node 


Note: Can't use topological sort as it work for directed graph

Approach #1: Union find
1. Take array parent. Index represent the node and value represent the parent. Init each node as its self parent
2. Take count set to no of nodes
3. implement find x
   1. if parent[x] != x
      1. return find parent[x]  
   2. else 
      1. return x
4. implement union `x`, `y`
   1. if parent of x != parent of y
      1. set parent of y = parent of x
      2. decrease count--
5. Actual code
   1. for each edge 
      1. if find x == find y return false //`cycle detected because as they have a same parent and if x,y have connection, it will make a cycle`
      2. union x, y
6. return count == 1 //`means only one component return true` 
Approach #2: DFS (or BFS)
1. Visited bool array, preNode to store previous visited node
2. Run form any node
   1. DFS -> bool
      1. if node is in visited, return false
      2. Add to visited
      3. for each adj node 
         1. if adj node == preNode, continue 
         2. if dfs is not true return false 
      4. return true
3. if not dfs return true, check for connected by count true in visited and no of nodes 

-----

### 323: Number of connected components (find it on lint code)
Intuition:
1. Use DFS 
2. Union-find. 
   1. Fist union will the build the array of parents
   2. For each node, find the parent and return number of parents

Approach #1: Union-Find
1. Take array parent where , index represent the node and value represent the parent. To start, each value should be node
2. Take array rank, index represent the node. init with rank `0` for all the nodes
3. implement find (x)
   1. if parent[x] != x 
      1. parent[x] = find(`parent[x]`) //path compression
4. implement union (x, y)
   1. if parent[x] != parent [y]
      1. rank[p1] > rank[p2]
         1. parent[p2] = parent[p1]
      2. rank[p1] < rank[p2]
         1. parent[p1] = parent[p2]
      3. else
         1. parent[p2] = parent[p1]
         2. rank[p1]++
5. for each edge in edges
   1. run union
6. for each node, run find and return the unique parents
https://www.youtube.com/watch?v=gKKATlgNNqM


class Solution:
    """
    @param n: the number of vertices
    @param edges: the edges of undirected graph
    @return: the number of connected components
    """
    def count_components(self, n: int, edges: List[List[int]]) -> int:
        # write your code here
        parents = [i for i in range(n)]
        ranks = [0] * 5
        component = 5
        def find(x): 
         if parents[x] != x :
            parents[x] = find(parents[x])
         return parents[x]
        
        def union(x, y):
         p1 = parents[x]
         p2 = parents[y]
         if p1 != p2:
            r1 = ranks[x]
            r2 = ranks[y]
            if r1 > r2 :
               parents[p2] = parents[p1] 
            else if r2 > r1 
               parents[p1] = parents[p2]
            else
               parents[p2] = parents[p1] 
               rank[p1]++
               component--
         for each edge in edges:
            union(edge[0], edge[1])
         return component



Approach #2: DFS
1. Build adjacency list from the graph edges
2. Take visited bool array to the length of count of edges
3. for each node 
   1. if node is not in the visited
      1. increase the count 
      2. run dfs
         1. for each adjacency nodes 
            1. if adjacency node in not in visited, run dfs


----- 
### 207. Course Schedule
Intuition: 
1. prerequisites = [[1,0]] represent by 0 must be done first (before) and then 1. 0 -> 1. Build a directed graph in adjacency list where for each node, adjacency list will contains all the nodes which are dependent on the node  
2. Run the topological sort and try to detect the cycle. This can be done if node is already in visited 

Approach #1 DFS
1. Build adjacency list
2. Take visited and visiting bool array with the length of nodes
3. Visited will hold nodes true for which all the dfs is done
4. Visiting will hold current dfs path, and if node already added, means found the cycle
5. for each node, run dfs
   1. if dfs return false, return false
6. in dfs
   1. if node already in visiting, cycle detected, return false 
   2. if node in visited, return true
   3. mark visiting true
   4. for each adjacency node
      1. if not visited, run dfs
   5. mark visited true

----- 

### 417. Pacific Atlantic Water Flow
Intuition: Find water flow towards Pacific and Atlantic separately. Overlapping values are going to flow for both the ocean

Approach #1: dfs
1. create two bool array for visited track (one for pacific and another for atlantic) of array of the same size as heights array
2. dfs function which take four parameters heights, curr pair, previous pair, visited 
3. Traverse top row & bottom row, call dfs  by passing pacific and atlantic 
4. Traverse left row & right row,  call dfs  by passing pacific and atlantic 
5. find common true from pacific and atlantic array and return

---- 

### 695. Max Area of Island
Intuition: Same a Number of island (#200) , here we have to count the individual land in a single island and return the max count

Approach #1: BFS
1. for each element of grid  i and j
   1. if `grid[i][j]` equal to `1`
      1. adding i, j pair into the queue, set `grid[i][j]` to zero as a visited, increase the currIsLandArea to 1 
      2. for `queue` is empty 
         1. pop the pair, for each offset `[[-1,0], [1,0], [0,-1], [0,1]]` check if new pair in boundary and `1`
         2. adding new pair into the queue, set `grid[i][j]` to zero as a visited, increase the currIsLandArea to 1 
      3. compare the maxArea 
2. return maxArea   

Approach #2: DFS

1. for each element of grid  i and j
   1. if `grid[i][j]` equal to `1`
      1. run dfs(grid, curr[i, j]) which will return the area
         1. for each offset `[[-1,0], [1,0], [0,-1], [0,1]]` check if new pair in boundary and `1`
            1. set `grid[i][j]` to zero as a visited, increase the currIsLandArea with recursive dfs return
         2. return currIsLandArea
      2. compare the maxArea 


---- 
### 133. Clone Graph

Intuition: Use DFS to traversal the nodes, if node already visited in map - copy its neighbor else do the dfs


1. take a array (check constrain for size). Create first clone node and add it to Val index in array
2. Call dfs `dfs(node, cloned)`
   1. for each neighbor of node
      1. if neighbor node is not in cloned array
         1. create a clone node of neighbor, put clone node into array with index of Val
         2. run the dfs on neighbor node `dfs(neighbor, cloned)`
      2. (no else) get the cloned node from cloned array via curr_node's Val, get cloned neighbor node valve from cloned array via  neighbor node's Val and append to cloned node's neighbor

Mistakes
//Didn't notice constrains which says `1 <= Node.val <= 100` this reduce the need to creating map, it can done via array of 101 size
// if cloned node in not in clone array, add into that. There is no need to else statement, adding neighbor for all 
//Naming mistakes
// range on slice, first return is index, second is stored value 


Approach #2 BFS
1. Take a clone node and array of 101 len (check constrain), node which are cloned keep adding into the array on Val index
2. Add visited node into the queue, run unit queue is empty
   1. pop for queue
   2. for each neighbor 
      1. if neighbor is not in array, add it neighbor into array, adding into queue
      2. (no else, for all) get the cloned neighbor from array and adding to cloned curr node (find in array by currNode.Val)




-----
### 200. Number of Islands


Intuition: traversal of lands (1s) in DFS of BFS 

Approach #1: DFS

1. Two inner loop to traversal the matrix
2. If value is 1 call DFS function and increase the count of landCount = landCount+1
3. call `dfs(i, j)` 
   1. check for i and j are within boundary of grid along with if `grid[i][j]` is 1 else return (base case for recursive dfs call)
   2. mark `grid[i][j]` to `0` means visited
   3. now move to all four direction recursively 
   4. `dfs(i-1, j)`, `dfs(i, j-1)`, `dfs(i+1, j)`, `dfs(i, j+1)`
4. return `landCount`

Mistakes
//modifying grid, pass as a pointer
//when try to change or get (len or value) dereference `(*grid)[i][j]==1`
//Pass grid to recursive function, pass as it is
//byte use '' instead of ""

Approach #2: BFS
1. Two inner loop to traversal the matrix
2. If value is 1, increase landCount by 1
3. add i and j pair into queue
4. Mark visited by setting `grid[i][j]=0`
5. while queue is empty
   1. dqueue pair i, j from queue
   2. change i and j to travel in all four direction 
   3. check if they are in boundary of grid and value is 1
      1. add into the queue
      2. Mark visited by setting `grid[i][j]=0`
6. return land count




------
### 78. Subsets

1. Intuition: Need to find all possible subset (super-set/the power set), backtracking 
2. if start index >= len(nums) : add path into result and return
3. Choice : all the numbers in the unique array
4. Choose : add one element at a time on path (this can be done by in recursive call by current choice index + 1 )
5. Explore : recursive cal on above choose and next choice
6. Un-choose : remove the chosen element from path
7. Explore : recursive cal for next choice

--------
### 39. Combination Sum
Need to find all possible combination of given target where numbers can be repetitive
1. Intuition: Need to find all combination - hence backtracking.
   1. Same as 78. Subset with once change, instead of choosing new number each time, chose same number until, 
2. if start index >= len(nums) and target == sum of path: add path into result and return
3. if index >= len(nums) OR target > sum of path: return, it means that path does't contains the sum
4. Choice: all the numbers 
5. Choose: One element until either target == sum of path
6. Explore: recursive call for same start point
7. Un-choose: remove the added element from path
8. Explore: recursive call on next number 



----------
### 215. Kth Largest Element in an Array
1. Intuition #1: Kth largest, means a min heap of len k which contains elements greater then kth. This can be done as same in #703
2. Intuition #2: Sorting and heap solution will give n*log(n) time complexity. This can be done in n time complexity via quick select
3. Quick select: to find kth smallest element. 
   1. Need to find largest k, so update k = len of array - k //*imp*
   2. Partition: Goal, place random chosen pivot element to its true index
   3. if k is greater then pivot index, run quick select on left portion of partition else on right partition
   4. if kth value is equal to pivot, found the kth smallest element
   5. Partition logic
      1. take a random pivot and place it at right side of array by swap
      2. take a variable `pivotLessElementsTailIndex` and set to zero. This variable indicate all the element left to this are less then pivot.
      3. Loop through the array
         1. check if current index element is less then pivot, it means this element must be left side of `pivotLessElementsTailIndex`.
         2. swap current index element with `pivotLessElementsTailIndex` and `pivotLessElementsTailIndex` increment by 1.
      4. Once loop ends, swap `pivotLessElementsTailIndex` and right most element. This step will place pivot at right index where all the less then pivot are left and greater then are on right
      5. return the pivot index
   6. *Imp* Quick select can be implement as a recursive as well as iterative

-----------
### 973. K Closest Points to Origin
1. Intuition: Need to find closest point, so max heap of distance of length k. After push, if len of heap is greater than k, pop the max. This will make sure heap will have kth smallest distance 
2. Max heap of points 
3. Less func should calculate the distance and compare
4. add points one by one and check the the len of heap
5. if len of heap is greater then k, pop (which will pop the max distance point so far)
6. return the  heap  

---------
### 1046. Last Stone Weight
1. Intuition: max heap, and while len(heap) >1, keep pop two item and push the diff
 --------

### 703. Kth Largest Element in a Stream
1. Intuition: min heap with length k with only contains elements of largest then kth. Pop the kth largest element 
2. Stream require heap for find the min or max element as with every insert of new element, require to pull the min/max on top of the tree
3. create a type named `intHeap` of `[]int`
4. implement min heap functions on `intHeap`
5. Because kth largest element need to find, min heap need to implement of k length
6. if len(inHeap) is less than k, push to the heap
7. Else if element is greater than min element, pop, and push
8. Else ignore  

-----
### Heap in Go
1. Heap will be visually represent via tree
2. Heap will be stored via array

      100
      /  \
     19  36
     /\   /  \
   17 3  25   1
   /\          
  2  7

  [100, 19, 36, 17, 3,25, 1, 2, 7]
3. Heap in go require to implement container/heap interface
4. container/heap have following func which need to implement
   1. `Len() int` on value receiver
   2. `Less(i, j int) bool` on value receiver
   3. `Swap(i, j int)`  on value receiver
   4. `Push(x interface{})` on pointer receiver
   5. `Pop() interface{}` on pointer receiver
5. heap need to init by `heap.Init()`
6. Push and pop will be done via `heap.Push` and `heap.Pop` 

### 211. Design Add and Search Words Data Structure
1. Add word will be standard trie add
2. Search, need to support `.` as match all 
   1. For each char in word
      1. if char is `.`
         1. Need to search recursive for all all available child nodes for remaining characters from word
         2. if any of the search found the remaining element return true
         3. Return false if loop ends without returning the true
      2. else run standard search



### 1804 Implement Trie -II

1. Raw implementation will be same as regular trie 
2. Trie struct will have additional filed name CountWord, CountStartWith
3. When inserting the word into the trie, update CountStartWith for each character insert
4. Once each character is added into the trie, update the CountWord
5. While search, return appropriate count if word pr prefix exist 


### 208. Implement Trie (Prefix Tree)
Refer to DSA/trie   



### Recursive Sorting 
This is not going to most optimized solution for sorting, but help in understanding the sorting 
Given a nums array of unsorted, function should return the sorted array
1. Base case: array with single element will be sorted 
2. Input reduction: Run the sort function recursive for n-1 length array, before that store nth element 
3. Actual job: Insert the nth element in n-1 length array which is sorted on step 2
   1. To insert either can be done via iterative or recursive 
   2. Insert element function 
      1. base case: array with 0 element or last element > insert candidate
         1. insert element 
      2. Input reduction insert n-1, n
      3. Add value at the end of array 

-----

### 572. Subtree of Another Tree

Approach #1: Recursive DFS
1. if tree and sub-tree is null, return true
2. if one of the them is null, return false
3. Use the #100 Same tree function 
4. if isSame on root and sub-root is true
   return true
5. Else
   1. return isSubTree(root.left, subtree) OR isSubTree(root.left, subtree)  // if one of the true, we are good
   
------

### 100. Same Tree

Approach #1: Recursive DFS
1. if both the root are null return true
2. if one of the root is null return false
3. base case p.root.val == q.root.val
   1. return isSame(left of p & q) && isSame(right of p & q)  //both side need to be true

------

### 110. Balance of Binary Tree

Approach #1: Recursive DFS
1. Find depth/heigh of left and right sub tree 
2. Abs diff must greater than 1 return false
3. Find depth/height
   1. find the left depth/height
   2. fine the right depth/height
   3. return max of (left depth or right depth)+ 1 

-----

### 543. Diameter of Binary Tree

Approach #1: Recursive DFS
1. Its based on finding the depth/height of the binary tree
2. Diameter of binary tree will depth of left sub-tree + depth of right sub-tree
3. maxDepth out side of recursive depth function
4. find depth recursive  by
   1. find the left depth
   2. find the right depth
   3. sum left depth and right depth
   4. compare with maxDepth and change the maxDepth accordingly //trick
   5. return depth/heigh max of (left or right) + 1 //trick adding 1 include the current root node

----

### 104. Maximum Depth of Binary Tree

Approach #1 Recursive DFS
1. Base exit case if root is nil, return 0 as a height
2. Start from root, calculate the heigh of left sub-tree and the right sub-tree 
3. take the max of heights and add 1
4. repeat the process 

Approach #2 Iterative BFS

1. Use queue, add root, sel level 0
2. Iterative till len(queue) > 0
   1. Iterative for len of the current queue -  *trick, store len in tmp variable and run loop till the tmp len* 
      1. deque from top, if left node is not null, add in the queue and same for right node 
   2. increase the level
3. return the level

----
Approach #3 Iterative BFS using stack
1. Use stack, add root
2. use map to maintain the level, add key as root and value a level =1
3. take maxLevel = 1
4. Iterative till len(stack) >0
   1. pop all the top node from stack 
   2. Get the level of the node for map
      1. if pop-node.left != nil add into the stack and map with level+1
      2. if pop-node.right != nil add into the stack and map with level+1
   3. check id level > maxLevel, set maxLevel = level
5. return maxLevel


----
### 226. Invert Binary Tree

Approach #1
1. Check the exit case if root == nil
2. Store the root's right on the temp
3. set right = left and left = temp
4. run above recursive for root.right and root.left 

----



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

### 74. Search a 2D Matrix

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