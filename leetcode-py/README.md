# Python cheat sheets



`words: List[str]` create a map store graph as an adjacency list where chars become the the key and value becomes adjacency list/array 
 `graph = {char: [] for word in words for char in word}`
 this will create following map
 
 `input = ["wrt", "wrf", "er", "ett", "rftt"]`
 output `{'w': [], 'r': [], 't': [], 'f': [], 'e': []}`
 
 ----

Loop over list and get the index and value both
`for i, interval in enumerate(intervals):`


Loop over list from non zero starting point
`for j in range(i+1, len(intervals)):`

----
Sorting 
- `sorted()` return the new sorted list without modifying existing 
-  `my_list.sort()` modify existing list
-  Reverse sorting pass `reverse=True`
-  custom sorting pass `key=lambda x:x[0]` 
```
my_list = [(1, 'apple'), (3, 'banana'), (2, 'orange')]

# Sort based on the second element of each tuple (the fruit name)
sorted_list = sorted(my_list, key=lambda x: x[1])
print(sorted_list)
# Output: [(1, 'apple'), (3, 'banana'), (2, 'orange')]

# Sort based on the first element of each tuple (the number)
my_list.sort(key=lambda x: x[0])
print(my_list)
# Output: [(1, 'apple'), (2, 'orange'), (3, 'banana')]
```

-----
maps

- How to check keys in maps
`key in map`

- How to check values in map
`value in map.values()`

----
List

- How to define an array of 26 size with all the element `0`
`unique = [0] * 26`

