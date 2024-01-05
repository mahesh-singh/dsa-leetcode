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
        visited = [False] * n
        adjMap = {i:[] for i in range(n)}
        for edge in edges:
            adjMap[edge[0]].append(edge[1])
            adjMap[edge[1]].append(edge[0])
        
        def dfs(node, preNode):
            if visited[node]:
                return False

            visited[node] = True
            for adj in adjMap[node]:
                if adj == preNode:
                    continue
                if not dfs(adj, node): 
                    return False

            
            return True

        if not dfs(0, -1):
            return False
        
        return all(visited)
# [] internal function and self



