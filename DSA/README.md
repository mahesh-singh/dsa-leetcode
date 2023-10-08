### Big O
Big O - worst case 
Clasification  - 3:41

### Dynamic array

### Linked list 

#### Complexity
- Search
  - Singly O(n)
  - Doubly O(n)
- Insert at head
  - Singly O(1)
  - Doubly O(1)
- Insert at tail
  - Singly O(1) 
  - Doubly O(1) 
- Remove at head
  - Singly O(1)
  - Doubly O(1)
- Remove at tail
  - Singly O(n) : because need to find new tail
  - Doubly O(1) : because tail have previous pointer which will become the new tail
- Remove at middle
  - Singly O(n) 
  - Doubly O(n) 
 
 #### Linked list traversal
1. while 
2. recursive 

#### Linked list Tortoise and Hare 


### Time complexity 
1. `for i=0; i<n; i++` O(n)
2. `for i=0; i<n; i=i+2` O(n)
3. `for i=n; i>0; i=--` O(n)
4. `for i=0; i<n; i=i*2` O(log n) base 2
5. `for i=0; i<n; i=i*3` O(log n) base 3
6. `for i=0; i<n; i=i/2` O(log n) base 2

Watch till 1.11

### Binary Tree & Binary Search tree
1. Binary tree have at most two children
2. Binary search tree: Left sub tree has smaller and right sub tree have larger elements
3. examples 
```
     5
    / \
  9     8
/       /\
1      7  10
      /\
    6   9

```

Above tree is *not a BST* as 9 is in the left subtree of 8

-------

```
    1
      \
      20
      /
    19
    /
  2
    \
      3
        \
          18
        /
      17

```
Above tree is BST as each node's right sub tree have larger element and left have smaller element 