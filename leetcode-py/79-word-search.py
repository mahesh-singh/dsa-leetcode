class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        ROW = len(board)
        COL = len(board[0])

        def dfs(i, j, word):
            #End of the word, hence true
            if len(word) == 0:
                return True
            # out of the boundary or char is not exist at i & j
            if i<0 or i>=ROW or j<0 or j>=COL or board[i][j] != word[0]:
                return False
            
            # store in tmp and will be use in reset for next iteration
            tmp = board[i][j] # chose

            # mark visited for current visit
            board[i][j] = "#"

            #explore
            res = dfs(i+1, j, word[1:]) or dfs(i-1, j, word[1:]) or dfs(i, j+1, word[1:]) or dfs(i, j-1, word[1:])
    
            #un-chose
            board[i][j] = tmp
            return res
        
        for i in range(ROW):
            for j in range(COL):
                if dfs(i, j, word): #choice 
                    return True
        
        return False