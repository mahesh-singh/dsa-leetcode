class Solution:
    def partitionString(self, s: str) -> int:
        unique = {}
        ret = 1
        for i in range(len(s)):
            if s[i] in unique:
                ret +=1
                unique = {}
            unique[s[i]] = i

        return ret 