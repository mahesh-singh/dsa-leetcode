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
        # write your code here
        visited = [false] * n
        adjMap = {i:[] for i in range(n)}
        for edge range edges:
            adjMap[edge[0]].append(edge[1])
            adjMap[edge[1]].append(edge[0])

        def dfs(, ):
            if 


# [] internal function and self