class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        result  = nums[0]

        currMax, curMin = 1,1
        for num in nums:
            tmp = currMax * num
            currMax = max(currMax * num, curMin * num, num)
            curMin = min(tmp, curMin * num, num)
            result = max(currMax, result)
        
        return result