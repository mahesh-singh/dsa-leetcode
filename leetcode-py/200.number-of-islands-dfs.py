from typing import List

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if len(grid) ==0 or len(grid[0]) == 0:
            return 0

        islands_count = 0

        def dfs(i, j):
            #base base: out of range or already visited. Notice OR
            if i<0 or i>=len(grid) or j<0 or j>=len(grid[0]) or grid[i][j] == "0":
                return


            grid[i][j] = "0" # mark visited

            dfs(i-1, j) #above
            dfs(i+1, j) #below
            dfs(i, j-1) #left
            dfs(i, j+1) #right


        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == "1": # only non-visited
                    islands_count +=1
                    dfs(i, j)

        return islands_count
