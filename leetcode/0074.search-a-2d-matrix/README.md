# Sudo code

m x n matrix contain m rows and n column

Approach #1

i = 0 //rows
j = n-1 //column

Go to the first row last column
Upper bound `0 <i < m`, `0 < j < n`

```
while i > 0 && i< m  && j > 0 && j < n
    if  matrix[i][n-1] == target
        return true
    
    if `target`  > matrix[i][n-1]
        // move to the next row    
        i++
    else 
        // move to previous column
        j--
``` 
Complexity = O(m+n)


Approach #2

```
// Binary search on rows to find probable row index  

top = 0
bot = len(matrix)-1
row = 0
while top < bot
    row = (top + bot)/2

    if matrix[row][n-1] < `target`
        bot = row
    else matrix[row][0] > `target`
        top = row
    else
        break //already on probable row


binary search on the row `matrix[row]`

```

1. Binary search on rows to find probable row index  
2. if `target` > mid row last element; change change the top to mid
3. if `target` < mid row fist element; change the bottom top to mid
4. Otherwise you are on the probable row; break the loop   
5. Binary search on the the elements of the row

Complexity  = log(m) + log(n)   




