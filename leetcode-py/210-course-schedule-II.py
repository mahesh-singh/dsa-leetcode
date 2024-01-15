
class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        prereq = {c:[] for c in range(numCourses)}

        for course, pre in prerequisites:
            prereq[course].append(pre)
        
        visited = []
        visiting = [False]*numCourses

        def dfs(curr):
            if visiting[curr]:
                return False
            if curr in visited:
                return True
            
            visiting[curr] = True
            for adj in prereq.get(curr, []):
                if not dfs(adj):
                    return False
            
            visiting[curr]= False
            visited.append(curr)
            return True
        
        for c in range(numCourses):
            if not dfs(c):
                return []
        
        return visited