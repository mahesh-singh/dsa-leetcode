from typing import List
class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1
        curr_min = float('inf')
        while l < r:
            m =  l + (r-l)//2
            curr_min = min(curr_min, nums[m])

            if nums[m] > nums[r]:
                # move right
                l = m+1
            else:
                r = m-1

        return min(curr_min, nums[l])
