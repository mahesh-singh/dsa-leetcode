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
    components = 5

    def find(x):
      if parents[x] != x:
        parents[x] = find(parents[x])
      return parents[x]

    def union(x, y):
      nonlocal components
      p1 = parents[x]
      p2 = parents[y]

      if p1 != p2:
        r1 = ranks[p1]
        r2 = ranks[p2]
        if r1 > r2:
          parents[p2] = parents[p1]
        elif r2 > r1:
          parents[p1] = parents[p2]
        else:
          parents[p2] = parents[p1]
          ranks[p1] += 1
        components -= 1
    for edge in edges:
      union(edge[0], edge[1])
    return components