class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        res = []
        last_index = {}
        for i in range(len(s)):
            last_index[s[i]] = i

        count = 0
        break_point = 0
        for i in range(len(s)):
            last = last_index[s[i]]
            break_point = max(last, break_point)
            count += 1

            if break_point == i:
                res.append(count)
                count = 0
        return res