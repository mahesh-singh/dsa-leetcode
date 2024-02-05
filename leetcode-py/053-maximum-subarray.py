
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxSum = float('-inf')
        curr_max = 0
        start, end, tmp_start = 0, 0, 0
        for i, n in enumerate(nums):
            curr_max = n+curr_max
            if curr_max > maxSum:
                maxSum = curr_max
                end = i
                start = tmp_start

            if curr_max < 0:
                curr_max =0
                tmp_start = i+1
        print(start, end)
        return maxSum
