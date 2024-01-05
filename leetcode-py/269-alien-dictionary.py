from typing import (
    List,
    OrderedDict, )


class Solution:
  """
    @param words: a list of words
    @return: a string which is correct order
    """

  def alien_order(self, words: List[str]) -> str:
    # Write your code here
    graph = {char:[] for word in words for char in word}

    
    
    for i in range(len(words) - 1):
      word1 = words[i]
      word2 = words[i + 1]
      minLength = min(len(word1), len(word2))
      if len(word1) > len(word2) and word1[:minLength] == word2[:minLength]:
        return ""
      for j in range(minLength):
        if word1[j] != word2[j]:
          graph[word1[j]].append(word2[j])
          break

    
    visited = []
    visiting = OrderedDict((char, False) for char in graph)
    print(visiting)
    
    def dfs(curr):
      if curr in visited:
        return True
      if visiting[curr]:
        return False

      visiting[curr] = True

      for adj in graph[curr]:
        if not dfs(adj):
          return False

      visiting[curr] = False
      print(curr)
      print(visited)
      visited.append(curr)
      print(visited)
      return True

    
    for key in graph:
      if not dfs(key):
        return ""
    print(visited)
    print(list(visited))
    return "".join(reversed(visited))