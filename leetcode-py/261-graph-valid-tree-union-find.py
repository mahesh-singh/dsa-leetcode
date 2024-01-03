from typing import (
    List,
)

class Solution:
    """
      @param n: An integer
      @param edges: a list of undirected edges
      @return: true if it's a valid tree, or false
      """
    def valid_tree(self, n: int, edges: List[List[int]]) -> bool:
        parents = [i for i in range(n)]

        components = n
        def find(x):
              
            if parents[x] != x:
                parents[x] = find(parents[x])
            return parents[x]

        def union(x, y):
            nonlocal components
            p1 = find(x)
            p2 = find(y)
            if p1 == p2:
                return False
            parents[p2] = parents[p1]
            components -= 1
            return True

        for u, v in edges:
            if not union(u, v):
                return False


        return components == 1
